
attribute serialization {
    skip: Bool
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
// Bar @alias('Foo')
//
// for field alias
// type Foo {
//   bar: Bar @alias('foobar')  
// }
attribute alias: String


/// make the integer being marshaled as signed integer in protobuf if set true
attribute signed: Bool

attribute disable_generate_go_format_codec: Bool

attribute disable_generate_go_union: Bool
attribute disable_generate_go_union_codec: Bool
attribute disable_generate_go_union_ext: Bool

attribute disable_auto_generation: Bool