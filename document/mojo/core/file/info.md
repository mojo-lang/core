| 字段 | 类型 | 格式类型 | 是否必须 | 默认值 | 说明 |
|---|---|---|---|---|---|
| `name` | `string` |  | 否 |  | base name of the file |
| `suffix` | `string` |  | 否 |  | suffix of the file name |
| `size` | `integer` | `int64` | 否 |  | length in bytes for regular files; system-dependent for others |
| `changeTime` | `string` | `DateTime` | 否 |  | 文件的元数据发生变化。比如权限，所有者等 |
| `modifyTime` | `string` | `DateTime` | 否 |  | 文件内容被修改的最后时间 |
