| 字段 | 类型 | 格式类型 | 是否必须 | 默认值 | 说明 |
|---|---|---|---|---|---|
| `name` | `string` |  | 否 |  | the name of the file, usually it is the full path of the file |
| `isDir` | `boolean` |  | 否 |  | whether is the directory or not |
| `mode` | `string` |  | 否 |  | the value must be const to "dir" |
| `info` | `mojo.core.File.Info` |  | 否 |  |
| `files` | `Array<mojo.core.File>` |  | 否 |  | sub files in this file if is a directory |
