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

var DeclTypeNames = map[int32]string{
	0: "unspecified",
	1: "type",
	2: "value",
	3: "function",
	4: "constructor",
	5: "attribute",
	6: "package",
	7: "generic",
}

var DeclTypeValues = map[string]DeclType{
	"unspecified": DeclType_DECL_TYPE_UNSPECIFIED,
	"type":        DeclType_DECL_TYPE_TYPE,
	"value":       DeclType_DECL_TYPE_VALUE,
	"function":    DeclType_DECL_TYPE_FUNCTION,
	"constructor": DeclType_DECL_TYPE_CONSTRUCTOR,
	"attribute":   DeclType_DECL_TYPE_ATTRIBUTE,
	"package":     DeclType_DECL_TYPE_PACKAGE,
	"generic":     DeclType_DECL_TYPE_GENERIC,
}

func (x DeclType) Format() string {
	v := int32(x)
	if s, ok := DeclTypeNames[v]; ok {
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

func (x DeclType) ToString() string {
	return x.Format()
}

func (x *DeclType) Parse(value string) error {
	if x != nil && len(value) > 0 {
		if s, ok := DeclTypeValues[value]; ok {
			*x = s
		} else {
			v := CaseStyler("snake")(value)
			if s, ok = DeclTypeValues[v]; ok {
				*x = s
			} else {
				return fmt.Errorf("invalid DeclType: %s", value)
			}
		}
	}
	return nil
}

func ParseDeclType(value string) (DeclType, error) {
	var v DeclType
	if err := (&v).Parse(value); err != nil {
		return v, err
	}
	return v, nil
}
