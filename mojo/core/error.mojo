/// # Error

/// ErrorCode is usually static and predefined in a specific domain
/// "system.121"
/// "system.NotSupport"
/// code `[100..599]` is reserved by duplicated with http status code
/// code `0` is also reserved because of it usually has success meaning
@template('[{domain}.]{value}')
type ErrorCode {
    ///
    value: Int @1 @required

    /// a brief description for the code,
    /// eg. `InvalidArgument`, which can be used for index.
    brief: String @2
    
    /// system, runtime, ...
    domain: String @3
    
    /// the api document url for this error code
    document: Uri @4
}

/// @lang('zh-CN')
///
type Error {
    type Code = ErrorCode

    code: Code @1 | String @2 | Int @3
    message: String @4

    causes: [Error] @10
}
