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
	"strconv"
)

var OrderingSortNames = map[int32]string{
	1: "asc",
	2: "desc",
}

var OrderingSortValues = map[string]Ordering_Sort{
	"asc":  Ordering_SORT_ASC,
	"desc": Ordering_SORT_DESC,
}

func (x Ordering_Sort) Format() string {
	s, ok := OrderingSortNames[int32(x)]
	if ok {
		return s
	}
	if int(x) == 0 {
		return "unspecified"
	}
	return strconv.Itoa(int(x))
}

func (x Ordering_Sort) ToString() string {
	return x.Format()
}

func (x *Ordering_Sort) Parse(value string) error {
	if x != nil {
		s, ok := OrderingSortValues[value]
		if ok {
			*x = s
		} else {
			*x = Ordering_SORT_ASC
		}
	} else {
		*x = Ordering_SORT_ASC
	}
	return nil
}

func ParseOrdering_Sort(value string) (Ordering_Sort, error) {
	var v Ordering_Sort
	if err := (&v).Parse(value); err != nil {
		return v, err
	}
	return v, nil
}
