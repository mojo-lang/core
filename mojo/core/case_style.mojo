
/// Case Style types for name case style changing
/// 
/// for value, snake_case is the defualt style
/// for type, CamelCase is the default style
enum CaseStyle {
    unspecified @0
    snake @1           //< snake_case style
    screaming_snake @2 //< SCREAMING_SNAKE_CASE style
    kebab @3           //< kebab-case style
    screaming_kebab @4 //< SCREAMING-KEBAB-CASE style
    camel @5           // CamelCase style
    lower_camel @6     //< lowerCamelCase style
}
