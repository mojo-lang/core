| 字段 | 类型 | 格式类型 | 是否必须 | 默认值 | 说明 |
|---|---|---|---|---|---|
| `seconds` | `integer` | `int64` | 否 |  | Signed seconds of the span of time. Must be from -315,576,000,000 to +315,576,000,000 inclusive. |
| `nanoseconds` | `integer` | `int32` | 否 |  | Signed fractions of a second at nanosecond resolution of the span of time. Durations less than one secondare represented with a 0 seconds field and a positive or negative nanoseconds field. For durations of one second or more,a non-zero value for the nanoseconds field must be of the same sign as the seconds field. Must befrom -999,999,999 to +999,999,999 inclusive. |
