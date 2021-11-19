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

var DayOfWeekNames = map[int32]string{
	1: "Monday",
	2: "Tuesday",
	3: "Wednesday",
	4: "Thursday",
	5: "Friday",
	6: "Saturday",
	7: "Sunday",
}

var DayOfWeekValues = map[string]DayOfWeek{
	"Monday":    DayOfWeek_DAY_OF_WEEK_MONDAY,
	"Tuesday":   DayOfWeek_DAY_OF_WEEK_TUESDAY,
	"Wednesday": DayOfWeek_DAY_OF_WEEK_WEDNESDAY,
	"Thursday":  DayOfWeek_DAY_OF_WEEK_THURSDAY,
	"Friday":    DayOfWeek_DAY_OF_WEEK_FRIDAY,
	"Saturday":  DayOfWeek_DAY_OF_WEEK_SATURDAY,
	"Sunday":    DayOfWeek_DAY_OF_WEEK_SUNDAY,
}

func (x DayOfWeek) Format() string {
	s, ok := DayOfWeekNames[int32(x)]
	if ok {
		return s
	}
	if int(x) == 0 {
		return "unspecified"
	}
	return strconv.Itoa(int(x))
}

func (x *DayOfWeek) Parse(value string) {
	if x != nil {
		s, ok := DayOfWeekValues[value]
		if ok {
			*x = s
		} else {
			*x = DayOfWeek_DAY_OF_WEEK_MONDAY
		}
	} else {
		*x = DayOfWeek_DAY_OF_WEEK_MONDAY
	}
}
