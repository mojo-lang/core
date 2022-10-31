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

var WildcardStringWildcardTypeNames = map[int32]string{
	0: "unspecified",
	1: "single_character",
	2: "zero_or_more_characters",
	3: "in_characters",
	4: "not_in_characters",
}

var WildcardStringWildcardTypeValues = map[string]WildcardString_Wildcard_Type{
	"unspecified":             WildcardString_Wildcard_TYPE_UNSPECIFIED,
	"single_character":        WildcardString_Wildcard_TYPE_SINGLE_CHARACTER,
	"zero_or_more_characters": WildcardString_Wildcard_TYPE_ZERO_OR_MORE_CHARACTERS,
	"in_characters":           WildcardString_Wildcard_TYPE_IN_CHARACTERS,
	"not_in_characters":       WildcardString_Wildcard_TYPE_NOT_IN_CHARACTERS,
}

func (x WildcardString_Wildcard_Type) Format() string {
	v := int32(x)
	if s, ok := WildcardStringWildcardTypeNames[v]; ok {
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

func (x WildcardString_Wildcard_Type) ToString() string {
	return x.Format()
}

func (x *WildcardString_Wildcard_Type) Parse(value string) error {
	if x != nil && len(value) > 0 {
		if s, ok := WildcardStringWildcardTypeValues[value]; ok {
			*x = s
		} else {
			v := CaseStyler("snake")(value)
			if s, ok = WildcardStringWildcardTypeValues[v]; ok {
				*x = s
			} else {
				return fmt.Errorf("invalid WildcardString_Wildcard_Type: %s", value)
			}
		}
	}
	return nil
}

func ParseWildcardString_Wildcard_Type(value string) (WildcardString_Wildcard_Type, error) {
	var v WildcardString_Wildcard_Type
	if err := (&v).Parse(value); err != nil {
		return v, err
	}
	return v, nil
}
