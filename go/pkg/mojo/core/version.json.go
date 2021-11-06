package core

import (
	jsoniter "github.com/json-iterator/go"
	"strconv"
	"strings"
	"unsafe"
)

func init() {
	jsoniter.RegisterTypeDecoder("core.Version", &VersionCodec{})
	jsoniter.RegisterTypeEncoder("core.Version", &VersionCodec{})
}

type VersionCodec struct {
}

func (codec *VersionCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	value := iter.ReadString()
	version := (*Version)(ptr)

	if len(value) > 0 {
		segments := strings.Split(value, "+")
		if len(segments) > 1 {
			version.Builds = strings.Split(segments[1], ".")
		}

		segments = strings.Split(segments[0], "-")
		if len(segments) > 1 {
			version.PreReleases = strings.Split(segments[1], ".")
		}

		segments = strings.Split(segments[0], ".")
		if len(segments) == 3 {
			major, err := strconv.Atoi(segments[0])
			if err != nil {
				iter.ReportError("VersionCodec.DecodeMajor", err.Error())
			} else {
				version.Major = uint64(major)
			}

			minor, err := strconv.Atoi(segments[1])
			if err != nil {
				iter.ReportError("VersionCodec.DecodeMinor", err.Error())
			} else {
				version.Minor = uint64(minor)
			}

			patch, err := strconv.Atoi(segments[2])
			if err != nil {
				iter.ReportError("VersionCodec.DecodePatch", err.Error())
			} else {
				version.Patch = uint64(patch)
			}
		}
	}
}

func (codec *VersionCodec) IsEmpty(ptr unsafe.Pointer) bool {
	version := (*Version)(ptr)
	return version == nil || (version.Major == 0 && version.Minor == 0 && version.Patch == 0)
}

func (codec *VersionCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	version := (*Version)(ptr)
	stream.WriteUint64(version.Major)
	stream.WriteRaw(".")
	stream.WriteUint64(version.Minor)
	stream.WriteRaw(".")
	stream.WriteUint64(version.Patch)

	if len(version.PreReleases) > 0 {
		stream.WriteInt8('-')
		stream.WriteString(strings.Join(version.PreReleases, "."))
	}

	if len(version.Builds) > 0 {
		stream.WriteInt8('+')
		stream.WriteString(strings.Join(version.PreReleases, "."))
	}
}
