/// # Error

/// ErrorCode is usually static and predefined in a specific domain
/// "system.121"
/// "system.NotSupport"
/// code `[100..599]` is reserved by duplicated with http status code
/// code `0` is also reserved because of it usually has success meaning
@template('{domain.}{value}{.brief}')
@serialization_with_string
type ErrorCode {
    ///
    value: Int @1 @required

    /// system, runtime, ...
    domain: String @3

    /// a brief description for the code,
    /// eg. `InvalidArgument`, which can be used for index.
    brief: String @2
    
    /// a detail description for the code
    description: String @5
    
    /// the api document url for this error code
    document: Url @4
}

/// @lang('zh-CN')
///
type Error {
    ///
    code: ErrorCode @1
    
    ///
    message: String @4

    /// the api document url for this error
    document: Url @5

    // the structural information for the error
    // the parameters check: { expect: {a: '>0'}  actual: {a: -2} }
    //context 

    /// the detail errors which cause this error
    causes: [Error] @10

    //backtraces: [Backtrace]
}
