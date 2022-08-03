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
	"fmt"
	"strconv"
	"strings"
)

var FileModeNames = map[int32]string{
	0:  "unspecified",
	10: "dir",
}

var FileModeValues = map[string]File_Mode{
	"unspecified": File_MODE_UNSPECIFIED,
	"dir":         File_MODE_DIR,
}

func (x File_Mode) Format() string {
	v := int32(x)
	if s, ok := FileModeNames[v]; ok {
		if v == 0 && "unspecified" == strings.ToLower(s) {
			return ""
		}
		return s
	}
	if v == 0 {
		return ""
	}
	return strconv.Itoa(int(v))
}

func (x File_Mode) ToString() string {
	return x.Format()
}

func (x *File_Mode) Parse(value string) error {
	if x != nil && len(value) > 0 {
		if s, ok := FileModeValues[value]; ok {
			*x = s
		} else {
			v := CaseStyler("snake")(value)
			if s, ok = FileModeValues[v]; ok {
				*x = s
			} else {
				return fmt.Errorf("invalid File_Mode: %s", value)
			}
		}
	}
	return nil
}

func ParseFile_Mode(value string) (File_Mode, error) {
	var v File_Mode
	if err := (&v).Parse(value); err != nil {
		return v, err
	}
	return v, nil
}
