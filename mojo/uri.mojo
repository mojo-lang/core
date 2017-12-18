/// # URI
/// 
/// Uniform Resource Identifier (URI) is a string of characters used to identify a resource.
///
/// A Uniform Resource Name (URN) may be compared to a person's name,
/// while a Uniform Resource Locator (URL) may be compared to their street address.
/// In other words, a URN identifies an item and a URL provides a method for finding it.
///
/// https://tools.ietf.org/html/rfc3986
///
///
@template()
type Uri {
    @template('//[{user_info} '@'] {host} [":" {port}]')
    type Authority {
        user_info: String @1
        host: String @2 @format()
        port: Int @3
    }

    ///
    type Path {
        segments:[String]
    }

    ///
    type Query = [(String, String)]

    ///
    scheme: String! @1 @format(r'[\w\d+-.]+')
    
    ///
    authority: Authority @2
    
    ///
    path: Path! @3
    
    ///
    query: Query @4 | String @5
    
    ///
    fragment: String @6
}