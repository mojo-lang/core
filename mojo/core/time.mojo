/// # Time Module 
///

///
type Date {
    year:  Int32 @1
    month: Int32 @2
    day:   Int32 @3
}

///
type DateTime {
    year: Int32 @1
    month: Int32 @2
    day: Int32 @3

    hour: Int32 @4
    minute: Int32 @5
    seconds: Int64 @6
    nanoseconds: Int32 @7

    time_zone: TimeZone @15
}

/// represent timezone information in the locations where different offsets are used in different days of the year
/// or where historical changes have been made to civil time.
type TimeZone {
    /// representing the difference between the local time and UTC.
    /// It must be strictly between -24 * 60 * 60 and 24 * 60 * 60
    offset: Int32 @1 @signed

    /// 
    name:  String @2
}

///
type Timestamp {
    seconds: Int64 @1
    nanoseconds: Int32 @2
}

/// A Duration represents a signed, fixed-length span of time represented as a count of seconds and
/// fractions of seconds at nanosecond resolution. It is independent of any calendar and concepts like "day" or "month".
/// It is related to Timestamp in that the difference between two Timestamp values is a Duration and
/// it can be added or subtracted from a Timestamp. Range is approximately +-10,000 years.
type Duration {
    /// Signed seconds of the span of time. Must be from -315,576,000,000 to +315,576,000,000 inclusive.
    seconds: Int64 @1 @signed

    /// Signed fractions of a second at nanosecond resolution of the span of time. Durations less than one second
    /// are represented with a 0 seconds field and a positive or negative nanoseconds field. For durations of one second or more,
    /// a non-zero value for the nanoseconds field must be of the same sign as the seconds field. Must be
    /// from -999,999,999 to +999,999,999 inclusive.
    nanoseconds: Int32 @2 @signed
}

///
//func now() -> DateTime

///
//func local_now() -> DateTime

/// ## DateTime
/// 
///
//Timestamp(String)

///
//Timestamp(Date)

///
//func to<>(Timestamp) -> String

///
//func +(Timestamp, Duration) -> DateTime

///
//func -(Timestamp, Duration) -> DateTime

///
//func year(Timestamp) -> Int

///
//func month(Timestamp) -> Int @in([1..12])
