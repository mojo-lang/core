| 字段 | 类型 | 格式类型 | 是否必须 | 默认值 | 说明 |
|---|---|---|---|---|---|
| `subject` | `string` |  | 否 |  | The subject on which the quota check failed.For example, "clientip:" or "project:". |
| `description` | `string` |  | 否 |  | A description of how the quota check failed. Clients can use thisdescription to find more about the quota configuration in the service'spublic documentation, or find the relevant quota limit to adjust throughdeveloper console.<br>For example: "Service disabled" or "Daily Limit for read operationsexceeded". |
