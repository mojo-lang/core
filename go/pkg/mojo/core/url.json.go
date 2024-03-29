// Code generated by mojo. DO NOT EDIT.
// Rerunning mojo will overwrite this file.
//
// Copyright 2021 Mojo-lang.org
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package core

import (
	"unsafe"

	jsoniter "github.com/json-iterator/go"
)

func init() {
	RegisterJSONTypeDecoder("core.Url", &UrlStringCodec{})
	RegisterJSONTypeEncoder("core.Url", &UrlStringCodec{})
}

// BareUrl will be jsonify to raw, without any codec
type BareUrl Url

type UrlStringCodec struct {
	IsFieldPointer bool
}

func (codec *UrlStringCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	s := iter.ReadString()
	url := codec.url(ptr)
	if url == nil {
		url = &Url{}
		*(**Url)(ptr) = url
	}

	if err := url.Parse(s); err != nil {
		iter.ReportError("UrlStringCodec", err.Error())
	}
}

func (codec *UrlStringCodec) IsEmpty(ptr unsafe.Pointer) bool {
	url := codec.url(ptr)
	if url != nil {
		if checker, ok := interface{}(url).(EmptyChecker); ok {
			return checker.IsEmpty()
		}
		return false
	}
	return true
}

func (codec *UrlStringCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	url := codec.url(ptr)
	stream.WriteString(url.Format())
}

func (codec *UrlStringCodec) url(ptr unsafe.Pointer) *Url {
	if codec.IsFieldPointer {
		return *(**Url)(ptr)
	}
	return (*Url)(ptr)
}

type UrlStructCodec struct {
	IsFieldPointer bool
}

func (codec *UrlStructCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	url := codec.bareUrl(ptr)
	if a := iter.ReadAny(); a.ValueType() == jsoniter.ObjectValue {
		if url == nil {
			url = &BareUrl{}
			*(**BareUrl)(ptr) = url
		}
		a.ToVal(url)
	}
}

func (codec *UrlStructCodec) IsEmpty(ptr unsafe.Pointer) bool {
	url := (*Url)(codec.bareUrl(ptr))
	if url != nil {
		if checker, ok := interface{}(url).(EmptyChecker); ok {
			return checker.IsEmpty()
		}
		return false
	}
	return true
}

func (codec *UrlStructCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	stream.WriteVal(codec.bareUrl(ptr))
}

func (codec *UrlStructCodec) bareUrl(ptr unsafe.Pointer) *BareUrl {
	if codec.IsFieldPointer {
		return *(**BareUrl)(ptr)
	}
	return (*BareUrl)(ptr)
}
