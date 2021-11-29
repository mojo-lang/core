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

/// Wrapper type for the Generic T type
///
///
type Boxed<T> {
    val: T @1 //< the boxed value
}

/// Wrapper type for `Bool`
///
/// The JSON representation for `BoxedBool` is JSON boolean.
type BoxedBool = Boxed<Bool>

/// Wrapper type for `Int64`
///
/// The JSON representation for `BoxedInt64` is JSON number.
type BoxedInt64 = Boxed<Int64>

/// Wrapper type for `UInt64`
///
/// The JSON representation for `BoxedUInt64` is JSON number.
type BoxedUInt64 = Boxed<UInt64>

/// Wrapper type for `Int32`
///
/// The JSON representation for `BoxedInt32` is JSON number.
type BoxedInt32 = Boxed<Int32>

/// Wrapper type for `UInt32`
///
/// The JSON representation for `BoxedUInt32` is JSON number.
type BoxedUInt32 = Boxed<UInt32>

/// Wrapper type for `Int32`
///
/// The JSON representation for `BoxedFloat` is JSON number.
type BoxedFloat = Boxed<Float>

/// Wrapper type for `UInt32`
///
/// The JSON representation for `BoxedDouble` is JSON number.
type BoxedDouble = Boxed<Double>

/// Wrapper type for `String`
///
/// The JSON representation for `BoxedString` is JSON string.
type BoxedString = Boxed<String>

/// Wrapper type for `Bytes`
///
/// The JSON representation for `BoxedBytes` is JSON string.
type BoxedBytes = Boxed<Bytes>


/// Wrapper type for `Array<String>`
///
/// The JSON representation for `Int32s` is JSON boolean array.
type BoolValues: [Bool]

/// Wrapper type for `Array<String>`
///
/// The JSON representation for `Int32s` is JSON number array.
type Int32Values: [Int32]

/// Wrapper type for `Array<String>`
///
/// The JSON representation for `UInt32s` is JSON number array.
type UInt32Values: [UInt32]

/// Wrapper type for `Array<String>`
///
/// The JSON representation for `Int64s` is JSON number array.
type Int64Values: [Int64]

/// Alias type for `Int64s`
///
/// The JSON representation for `Integers` is JSON number array.
type IntegerValues = Int64s

/// Wrapper type for `Array<String>`
///
/// The JSON representation for `UInt64s` is JSON number array.
type UInt64Values: [UInt64]

/// Wrapper type for `Array<String>`
///
/// The JSON representation for `Floats` is JSON number array.
type FloatValues: [Float]

/// Wrapper type for `Array<String>`
///
/// The JSON representation for `Doubles` is JSON number array.
type DoubleValues: [Double]

/// Wrapper type for `Array<String>`
///
/// The JSON representation for `Doubles` is JSON string array.
type Strings: [String]
