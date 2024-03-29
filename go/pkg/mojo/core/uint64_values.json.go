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
	RegisterJSONTypeDecoder("core.UInt64Values", &UInt64ValuesCodec{})
	RegisterJSONTypeEncoder("core.UInt64Values", &UInt64ValuesCodec{})
}

type UInt64ValuesCodec struct {
}

func (codec *UInt64ValuesCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	any := iter.ReadAny()
	uint64Values := (*UInt64Values)(ptr)
	if any.ValueType() == jsoniter.ArrayValue {
		any.ToVal(&uint64Values.Vals)
	}
}

func (codec *UInt64ValuesCodec) IsEmpty(ptr unsafe.Pointer) bool {
	uint64Values := (*UInt64Values)(ptr)
	return uint64Values == nil || len(uint64Values.Vals) == 0
}

func (codec *UInt64ValuesCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	uint64Values := (*UInt64Values)(ptr)
	stream.WriteVal(uint64Values.Vals)
}
