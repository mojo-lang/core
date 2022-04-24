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

///
attribute max_length: Int64

///
attribute min_length: Int64

/// make sure the count of the container type's values is equal to the fixed_length
attribute fixed_length: Int64

/// make sure the number is less then and equal to the maximun (le <=)
@target(DeclType.value)
attribute maximum: Int64

///  (lt <)
attribute exclusive_maximum: Int64

///  (ge >=)
attribute minimum: Int64

///  (gt >)
attribute exclusive_minmum: Int64

/// (ne !=)
attribute not_equal: Int64

/// The multiple_of attribute can be used to limit numbers to multiples of a given number. It can be set to any positive number.
/// # example
/// ```
/// type Value: Int32 @multiple_of(10) // 0  OK, 10 OK, 20 OK, 23 NOT OK (not the multiple of 10)
/// ```
attribute multiple_of: UInt64


attribute unique: Bool

attribute nonempty: Bool

attribute precision: Int32