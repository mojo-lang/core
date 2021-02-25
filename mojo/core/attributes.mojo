
attribute serialization {
    skip: Bool
}

//attribute marshal {
//    skip: Bool
//}


/// attribute for the operator functions
attribute {

    ///
    prefix: Bool

    ///
    postfix: Bool

    ///
    infix: Bool

    ///
    literal: Bool
}

@infix
func operator +(left: String, right: String) -> String


/// r'.*'
@literal
@prefix
func operator r(str: String) -> Regex


// 5ms
@literal
@postfix
func operator ms(number: Int) -> Duration

// for type alias
// type Foo = Bar
// Bar @alias('Foo')
//
// for field alias
// type Foo {
//   bar: Bar @alias('foobar')  
// }
attribute alias: String