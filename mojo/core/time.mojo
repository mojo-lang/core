/// # Time Module 
///

///
type Date {}

///
type Timestamp {}

///
type Duration {}

///
func now() -> DateTime

///
func local_now() -> DateTime

/// ## DateTime
/// 
///
Timestamp(String)

///
Timestamp(Date)

///
func to<>(Timestamp) -> String

///
func +(Timestamp, Duration) -> DateTime

///
func -(Timestamp, Duration) -> DateTime

///
func year(Timestamp) -> Int

///
func month(Timestamp) -> Int @in([1..12])
