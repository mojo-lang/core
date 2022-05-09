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
	RegisterJSONTypeDecoder("core.Int32Values", &Int32ValuesCodec{})
	RegisterJSONTypeEncoder("core.Int32Values", &Int32ValuesCodec{})
}

type Int32ValuesCodec struct {
}

func (codec *Int32ValuesCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	any := iter.ReadAny()
	int32Values := (*Int32Values)(ptr)
	if any.ValueType() == jsoniter.ArrayValue {
		any.ToVal(&int32Values.Vals)
	}
}

func (codec *Int32ValuesCodec) IsEmpty(ptr unsafe.Pointer) bool {
	int32Values := (*Int32Values)(ptr)
	return int32Values == nil || len(int32Values.Vals) == 0
}

func (codec *Int32ValuesCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	int32Values := (*Int32Values)(ptr)
	stream.WriteVal(int32Values.Vals)
}
