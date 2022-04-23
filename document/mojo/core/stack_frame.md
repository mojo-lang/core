| 字段 | 类型 | 格式类型 | 是否必须 | 默认值 | 说明 |
|---|---|---|---|---|---|
| `instruction` | `integer` | `int64` | 是 |  | The program counter location as an absolute virtual address.  For theinnermost called frame in a stack, this will be an exact program counteror instruction pointer value.  For all other frames, this will be withinthe instruction that caused execution to branch to a called function,but may not necessarily point to the exact beginning of that instruction. |
| `module` | `mojo.core.CodeModule` |  | 否 |  | The module in which the instruction resides. |
| `functionName` | `string` |  | 否 |  | The function name, may be omitted if debug symbols are not available. |
| `functionBase` | `integer` | `int64` | 否 |  | The start address of the function, may be omitted if debug symbolsare not available. |
| `sourceFileName` | `string` |  | 否 |  | The source file name, may be omitted if debug symbols are not available. |
| `sourceLine` | `integer` | `int32` | 否 |  | The (1-based) source line number, may be omitted if debug symbols arenot available. |
| `sourceLineBase` | `integer` | `int64` | 否 |  | The start address of the source line, may be omitted if debug symbolsare not available. |
| `trust` | `string` |  | 否 |  | Amount of trust the stack walker has in the instruction pointerof this frame. |
