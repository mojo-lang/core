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

var StackFrameTrustNames = map[int32]string{
	0: "unspecified",
	1: "scan",
	2: "cfi_scan",
	3: "fp",
	4: "cfi",
	5: "prewalked",
	6: "context",
}

var StackFrameTrustValues = map[string]StackFrame_Trust{
	"unspecified": StackFrame_TRUST_UNSPECIFIED,
	"scan":        StackFrame_TRUST_SCAN,
	"cfi_scan":    StackFrame_TRUST_CFI_SCAN,
	"fp":          StackFrame_TRUST_FP,
	"cfi":         StackFrame_TRUST_CFI,
	"prewalked":   StackFrame_TRUST_PREWALKED,
	"context":     StackFrame_TRUST_CONTEXT,
}

func (x StackFrame_Trust) Format() string {
	v := int32(x)
	if s, ok := StackFrameTrustNames[v]; ok {
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

func (x StackFrame_Trust) ToString() string {
	return x.Format()
}

func (x *StackFrame_Trust) Parse(value string) error {
	if x != nil && len(value) > 0 {
		if s, ok := StackFrameTrustValues[value]; ok {
			*x = s
		} else {
			v := CaseStyler("snake")(value)
			if s, ok = StackFrameTrustValues[v]; ok {
				*x = s
			} else {
				return fmt.Errorf("invalid StackFrame_Trust: %s", value)
			}
		}
	}
	return nil
}

func ParseStackFrame_Trust(value string) (StackFrame_Trust, error) {
	var v StackFrame_Trust
	if err := (&v).Parse(value); err != nil {
		return v, err
	}
	return v, nil
}
