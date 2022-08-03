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

var ChecksumAlgorithmNames = map[int32]string{
	0: "UNSPECIFIED",
	1: "MD5",
	2: "SHA1",
	3: "SHA256",
	4: "SHA512",
}

var ChecksumAlgorithmValues = map[string]Checksum_Algorithm{
	"UNSPECIFIED": Checksum_ALGORITHM_UNSPECIFIED,
	"MD5":         Checksum_ALGORITHM_MD5,
	"SHA1":        Checksum_ALGORITHM_SHA1,
	"SHA256":      Checksum_ALGORITHM_SHA256,
	"SHA512":      Checksum_ALGORITHM_SHA512,
}

func (x Checksum_Algorithm) Format() string {
	v := int32(x)
	if s, ok := ChecksumAlgorithmNames[v]; ok {
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

func (x Checksum_Algorithm) ToString() string {
	return x.Format()
}

func (x *Checksum_Algorithm) Parse(value string) error {
	if x != nil && len(value) > 0 {
		if s, ok := ChecksumAlgorithmValues[value]; ok {
			*x = s
		} else {
			v := CaseStyler("screaming_snake")(value)
			if s, ok = ChecksumAlgorithmValues[v]; ok {
				*x = s
			} else {
				return fmt.Errorf("invalid Checksum_Algorithm: %s", value)
			}
		}
	}
	return nil
}

func ParseChecksum_Algorithm(value string) (Checksum_Algorithm, error) {
	var v Checksum_Algorithm
	if err := (&v).Parse(value); err != nil {
		return v, err
	}
	return v, nil
}
