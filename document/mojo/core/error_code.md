| 字段 | 类型 | 格式类型 | 是否必须 | 默认值 | 说明 |
|---|---|---|---|---|---|
| `val` | `integer` | `int32` | 是 |  |  |
| `name` | `string` |  | 否 |  | The name of the error code. This is a constant value that identifies theerror code. Error code name are unique within a particulardomain of errors. This should be at most 63 characters and match/[A-Z0-9_]+/. |
| `domain` | `string` |  | 否 |  | system, runtime, ... |
| `description` | `string` |  | 否 |  | a detail description for the code |
| `document` | `string` | `url` | 否 |  | the api document url for this error code |
| `httpStatusCode` | `integer` | `int32` | 否 |  | the http status code, which is error code mapping to |
