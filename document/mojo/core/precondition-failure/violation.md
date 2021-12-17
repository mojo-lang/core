| 字段 | 类型 | 格式类型 | 是否必须 | 默认值 | 说明 |
|---|---|---|---|---|---|
| `type` | `string` |  | 否 |  | The type of PreconditionFailure. We recommend using a service-specificenum type to define the supported precondition violation subjects. Forexample, "TOS" for "Terms of Service violation". |
| `subject` | `string` |  | 否 |  | The subject, relative to the type, that failed.For example, "google.com/cloud" relative to the "TOS" type would indicatewhich terms of service is being referenced. |
| `description` | `string` |  | 否 |  | A description of how the precondition failed. Developers can use thisdescription to understand how to fix the failure.<br>For example: "Terms of service not accepted". |
