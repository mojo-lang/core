
/// 对于type的format标注，说明该type类型序列化后的表现形式，并且该Type默认的json等序列化后的类型为String
/// 对于字段的format标注，说明该字段字符串需要符合的格式要求，一般为正则表达式
attribute format: TemplateString | Regex

