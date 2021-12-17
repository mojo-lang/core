| 字段 | 类型 | 格式类型 | 是否必须 | 默认值 | 说明 |
|---|---|---|---|---|---|
| `resourceType` | `string` |  | 否 |  | A name for the type of resource being accessed, e.g. "sql table","cloud storage bucket", "file", "Google calendar"; or the type URLof the resource: e.g. "type.googleapis.com/google.pubsub.v1.Topic". |
| `resourceName` | `string` |  | 否 |  | The name of the resource being accessed.  For example, a shared calendarname: "example.com_", if the currenterror is [google.rpc.Code.PERMISSION_DENIED][google.rpc.Code.PERMISSION_DENIED]. |
| `owner` | `string` |  | 否 |  | The owner of the resource (optional).For example, "user:" or "project:". |
| `description` | `string` |  | 否 |  | Describes what error is encountered when accessing this resource.For example, updating a cloud project may require the `writer` permissionon the developer console project. |
