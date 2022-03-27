package core

import (
    jsoniter "github.com/json-iterator/go"
    "unsafe"
)

func init() {
    jsoniter.RegisterTypeDecoder("core.Checksum", &ChecksumStringCodec{})
    jsoniter.RegisterTypeEncoder("core.Checksum", &ChecksumStringCodec{})
}

// BareChecksum will be jsonify to raw, without any codec
type BareChecksum Checksum

type ChecksumStringCodec struct {
    IsFieldPointer bool
}

func (codec *ChecksumStringCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
    value := iter.ReadString()
    checksum := codec.checksum(ptr)
    if checksum == nil {
        checksum = &Checksum{}
        *(**Checksum)(ptr) = checksum
    }

    if err := checksum.Parse(value); err != nil {
        iter.ReportError("ChecksumStringCodec", err.Error())
    }
}

func (codec *ChecksumStringCodec) IsEmpty(ptr unsafe.Pointer) bool {
    return codec.checksum(ptr).IsEmpty()
}

func (codec *ChecksumStringCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
    checksum := codec.checksum(ptr)
    stream.WriteString(checksum.Format())
}

func (codec *ChecksumStringCodec) checksum(ptr unsafe.Pointer) *Checksum {
    if codec.IsFieldPointer {
        return *(**Checksum)(ptr)
    }
    return (*Checksum)(ptr)
}

type ChecksumStructCodec struct {
    IsFieldPointer bool
}

func (codec *ChecksumStructCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
    checksum := codec.bareChecksum(ptr)
    if value := iter.ReadAny(); value.ValueType() == jsoniter.ObjectValue {
        if checksum == nil {
            checksum = &BareChecksum{}
            *(**BareChecksum)(ptr) = checksum
        }
        value.ToVal(checksum)
    }
}

func (codec *ChecksumStructCodec) IsEmpty(ptr unsafe.Pointer) bool {
    checksum := (*Checksum)(codec.bareChecksum(ptr))
    return checksum.IsEmpty()
}

func (codec *ChecksumStructCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
    c := codec.bareChecksum(ptr)
    stream.WriteVal(c)
}

func (codec *ChecksumStructCodec) bareChecksum(ptr unsafe.Pointer) *BareChecksum {
    if codec.IsFieldPointer {
        return *(**BareChecksum)(ptr)
    }
    return (*BareChecksum)(ptr)
}
