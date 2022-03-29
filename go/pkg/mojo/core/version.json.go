package core

import (
    jsoniter "github.com/json-iterator/go"
    "unsafe"
)

func init() {
    RegisterJSONTypeDecoder("core.Version", &VersionStringCodec{})
    RegisterJSONTypeEncoder("core.Version", &VersionStringCodec{})
}

// BareVersion will be jsonify to raw, without any codec
type BareVersion Version

type VersionStringCodec struct {
    IsFieldPointer bool
}

func (codec *VersionStringCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
    value := iter.ReadString()
    version := codec.version(ptr)
    if version == nil {
        version = &Version{}
        *(**Version)(ptr) = version
    }

    if err := version.Parse(value); err != nil {
        iter.ReportError("VersionStringCodec", err.Error())
    }
}

func (codec *VersionStringCodec) IsEmpty(ptr unsafe.Pointer) bool {
    return codec.version(ptr).IsEmpty()
}

func (codec *VersionStringCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
    version := codec.version(ptr)
    stream.WriteString(version.Format())
}

func (codec *VersionStringCodec) version(ptr unsafe.Pointer) *Version {
    if codec.IsFieldPointer {
        return *(**Version)(ptr)
    }
    return (*Version)(ptr)
}

type VersionStructCodec struct {
    IsFieldPointer bool
}

func (codec *VersionStructCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
    version := codec.bareVersion(ptr)
    if value := iter.ReadAny(); value.ValueType() == jsoniter.ObjectValue {
        if version == nil {
            version = &BareVersion{}
            *(**BareVersion)(ptr) = version
        }
        value.ToVal(version)
    }
}

func (codec *VersionStructCodec) IsEmpty(ptr unsafe.Pointer) bool {
    version := codec.bareVersion(ptr)
    return version == nil || (version.Major == 0 && version.Minor == 0 && version.Patch == 0)
}

func (codec *VersionStructCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
    stream.WriteVal(codec.bareVersion(ptr))
}

func (codec *VersionStructCodec) bareVersion(ptr unsafe.Pointer) *BareVersion {
    if codec.IsFieldPointer {
        return *(**BareVersion)(ptr)
    }
    return (*BareVersion)(ptr)
}
