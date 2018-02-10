/// # Time Module 
///

///
type Date {}

///
type DateTime {}

///
type Duration {}

///
func now() -> DateTime

///
func local_now() -> DateTime

/// ## DateTime
/// 
///
DateTime(String)

///
DateTime(Date)

///
func to<>(DateTime) -> String

///
func +(DateTime, Duration) -> DateTime

///
func -(DateTime, Duration) -> DateTime

///
func year(DateTime) -> Int

///
func month(DateTime) -> Int @in([1..12])
