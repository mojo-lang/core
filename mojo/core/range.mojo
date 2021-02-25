///
///
///
type Range<T> {
    min: T @1
    max: T @2

    min_excluded: Bool @3
    max_excluded: Bool @4
}

// 12...45    12..45
// 12..<45    12..<45
// 12<..45    12<..45
// 12<.<45    12<..<45