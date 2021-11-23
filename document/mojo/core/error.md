| 字段 | 类型 | 格式类型 | 是否必须 | 默认值 | 说明 |
|---|---|---|---|---|---|
| `code` | `mojo.core.ErrorCode` | `` | 否 |  | The error code |
| `message` | `string` | `` | 否 |  | A developer-facing error message, which should be in English. Anyuser-facing error message should be localized and sent in the[google.rpc.Status.details][google.rpc.Status.details] field, or localized by the client. |
| `details` | `Array<mojo.core.Any>` | `` | 否 |  | A list of messages that carry the error details.  There is a common set ofmessage types for APIs to use. |
