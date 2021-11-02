///
///
///
type Boxed<T> {
    value: T @1
}

///
///
///
type BoxedBool = Boxed<Bool>

///
///
///
type BoxedInt = Boxed<Int>

///
///
///
type BoxedString = Boxed<String>

///
type Strings : [String]

///
type Integers : [Int64]

///
type Doubles : [Double]
