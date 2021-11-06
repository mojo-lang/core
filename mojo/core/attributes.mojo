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

attribute serialization {
    skip: Bool
    omitempty: Bool @default(true)
}

//attribute unique: Bool

//attribute nonempty: Bool

//attribute marshal {
//    skip: Bool
//}


// attribute for the operator functions
//attribute {

    //
    //prefix: Bool

    //
    //postfix: Bool

    //
    //infix: Bool

    //
    //literal: Bool
//}

//@infix
//func operator +(left: String, right: String) -> String


/// r'.*'
//@literal
//@prefix
//func operator r(str: String) -> Regex


// 5ms
//@literal
//@postfix
//func operator ms(number: Int) -> Duration

// for type alias
// type Foo = Bar
// Bar @type_alias('Foo')
attribute type_alias: String

// for field alias
// type Foo {
//   bar: Bar @alias('foobar')  
// }
attribute alias: String


/// make the integer being marshaled as signed integer in protobuf if set true
attribute signed: Bool

//attribute const: String

/// attribute for struct
/// for example:
/// ```
/// @disable_generate({"go": ["ext", "fmt", "json"]})
/// type Foo = Bar @1
///          | Car @2
/// ```
/// the key is language type, may be one of "all", "go", "java", "python"
/// and the value, may be some of the "ext", "fmt", "json" 
attribute disable_generate: {String: [Strings]}

/// 对于union，序列化成json后，使用何种字段来表现类型，默认为type，可以是 @type，或是 $type
attribute discriminator: String

/// 对于union，转换成protobuf时，每个union的类型对于的字段名，如果是`{}`,则是类型名snake化
/// 如果是其他的，则会给每个字段附加`_val`后缀
attribute label_format: String