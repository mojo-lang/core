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
/// 实例化名称规格， TValue  TValues  TValuesMap
type Boxed<T> {
    val: T @1 //< the boxed value
}

/// Associative Values for string key based map
type AssocValues<T> = {String: T}

/// Wrapper type for `Bool`
///
/// The JSON representation for `BoxedBool` is JSON boolean.
@deprecated
type BoxedBool = Boxed<Bool>

/// Wrapper type for `Bool`
///
/// The JSON representation for `BoolValue` is JSON boolean.
type BoolValue = Boxed<Bool>

/// Wrapper type for `Int64`
///
/// The JSON representation for `BoxedInt64` is JSON number.
@deprecated
type BoxedInt64 = Boxed<Int64>

/// Wrapper type for `Int64`
///
/// The JSON representation for `Int64Value` is JSON number.
type Int64Value = Boxed<Int64>

/// Wrapper type for `UInt64`
///
/// The JSON representation for `BoxedUInt64` is JSON number.
@deprecated
type BoxedUInt64 = Boxed<UInt64>

/// Wrapper type for `UInt64`
///
/// The JSON representation for `UInt64Value` is JSON number.
type UInt64Value = Boxed<UInt64>

/// Wrapper type for `Int32`
///
/// The JSON representation for `BoxedInt32` is JSON number.
@deprecated
type BoxedInt32 = Boxed<Int32>

/// Wrapper type for `Int32`
///
/// The JSON representation for `Int32Value` is JSON number.
type Int32Value = Boxed<Int32>

/// Wrapper type for `UInt32`
///
/// The JSON representation for `BoxedUInt32` is JSON number.
@deprecated
type BoxedUInt32 = Boxed<UInt32>

/// Wrapper type for `UInt32`
///
/// The JSON representation for `UInt32Value` is JSON number.
type UInt32Value = Boxed<UInt32>

/// Wrapper type for `Float`
///
/// The JSON representation for `BoxedFloat` is JSON number.
@deprecated
type BoxedFloat = Boxed<Float>

/// Wrapper type for `Float`
///
/// The JSON representation for `FloatValue` is JSON number.
type FloatValue = Boxed<Float>

/// Wrapper type for `Double`
///
/// The JSON representation for `BoxedDouble` is JSON number.
@deprecated
type BoxedDouble = Boxed<Double>

/// Wrapper type for `Double`
///
/// The JSON representation for `DoubleValue` is JSON number.
type DoubleValue = Boxed<Double>

/// Wrapper type for `String`
///
/// The JSON representation for `BoxedString` is JSON string.
@deprecated
type BoxedString = Boxed<String>

/// Wrapper type for `String`
///
/// The JSON representation for `StringValue` is JSON string.
type StringValue = Boxed<String>

/// Wrapper type for `Bytes`
///
/// The JSON representation for `BoxedBytes` is JSON string.
@deprecated
type BoxedBytes = Boxed<Bytes>

/// Wrapper type for `Bytes`
///
/// The JSON representation for `BytesValue` is JSON string.
type BytesValue = Boxed<Bytes>

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
type IntegerValues = Int64Values

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
@deprecated
type Strings: [String]

/// Wrapper type for `Array<String>`
///
/// The JSON representation for `Doubles` is JSON string array.
type StringValues: [String]