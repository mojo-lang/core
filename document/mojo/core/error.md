| 字段 | 类型 | 格式类型 | 是否必须 | 默认值 | 说明 |
|---|---|---|---|---|---|
| `message` | `string` |  | 否 |  |  |
| `document` | `string` | `url` | 否 |  | the api document url for this error |
| `causes` | `Array<mojo.core.Error>` |  | 否 |  | the detail errors which cause this error |
| `code` | `mojo.core.ErrorCode` |  | 否 |  |  |
