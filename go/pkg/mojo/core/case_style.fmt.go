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

var CaseStyleNames = map[int32]string{
	0: "unspecified",
	1: "snake",
	2: "screaming_snake",
	3: "kebab",
	4: "screaming_kebab",
	5: "camel",
	6: "lower_camel",
}

var CaseStyleValues = map[string]CaseStyle{
	"unspecified":     CaseStyle_CASE_STYLE_UNSPECIFIED,
	"snake":           CaseStyle_CASE_STYLE_SNAKE,
	"screaming_snake": CaseStyle_CASE_STYLE_SCREAMING_SNAKE,
	"kebab":           CaseStyle_CASE_STYLE_KEBAB,
	"screaming_kebab": CaseStyle_CASE_STYLE_SCREAMING_KEBAB,
	"camel":           CaseStyle_CASE_STYLE_CAMEL,
	"lower_camel":     CaseStyle_CASE_STYLE_LOWER_CAMEL,
}

func (x CaseStyle) Format() string {
	v := int32(x)
	if s, ok := CaseStyleNames[v]; ok {
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

func (x CaseStyle) ToString() string {
	return x.Format()
}

func (x *CaseStyle) Parse(value string) error {
	if x != nil && len(value) > 0 {
		if s, ok := CaseStyleValues[value]; ok {
			*x = s
		} else {
			v := CaseStyler("snake")(value)
			if s, ok = CaseStyleValues[v]; ok {
				*x = s
			} else {
				return fmt.Errorf("invalid CaseStyle: %s", value)
			}
		}
	}
	return nil
}

func ParseCaseStyle(value string) (CaseStyle, error) {
	var v CaseStyle
	if err := (&v).Parse(value); err != nil {
		return v, err
	}
	return v, nil
}
