package core

import (
    jsoniter "github.com/json-iterator/go"
    "unsafe"
)

func init() {
    RegisterJSONTypeDecoder("core.Platform", &PlatformStringCodec{})
    RegisterJSONTypeEncoder("core.Platform", &PlatformStringCodec{})
}

// BarePlatform will be jsonify to raw, without any codec
type BarePlatform Platform

type PlatformStringCodec struct {
    IsFieldPointer bool
}

func (codec *PlatformStringCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
    value := iter.ReadString()
    platform := codec.platform(ptr)
    if platform == nil {
        platform = &Platform{}
        *(**Platform)(ptr) = platform
    }

    if err := platform.Parse(value); err != nil {
        iter.ReportError("PlatformStringCodec", err.Error())
    }
}

func (codec *PlatformStringCodec) IsEmpty(ptr unsafe.Pointer) bool {
    platform := codec.platform(ptr)
    return platform == nil || (platform.Architecture == 0 && platform.Os == 0)
}

func (codec *PlatformStringCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
    platform := codec.platform(ptr)
    stream.WriteString(platform.Format())
}

func (codec *PlatformStringCodec) platform(ptr unsafe.Pointer) *Platform {
    if codec.IsFieldPointer {
        return *(**Platform)(ptr)
    }
    return (*Platform)(ptr)
}

type PlatformStructCodec struct {
    IsFieldPointer bool
}

func (codec *PlatformStructCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
    platform := codec.barePlatform(ptr)
    if value := iter.ReadAny(); value.ValueType() == jsoniter.ObjectValue {
        if platform == nil {
            platform = &BarePlatform{}
            *(**BarePlatform)(ptr) = platform
        }
        value.ToVal(platform)
    }
}

func (codec *PlatformStructCodec) IsEmpty(ptr unsafe.Pointer) bool {
    platform := codec.barePlatform(ptr)
    return platform == nil || (platform.Architecture == 0 && platform.Os == 0)
}

func (codec *PlatformStructCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
    stream.WriteVal(codec.barePlatform(ptr))
}

func (codec *PlatformStructCodec) barePlatform(ptr unsafe.Pointer) *BarePlatform {
    if codec.IsFieldPointer {
        return *(**BarePlatform)(ptr)
    }
    return (*BarePlatform)(ptr)
}
