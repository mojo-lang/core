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
@template("{scheme}://{authority}/{path}{#fragment}{?query}")
type Url {
    @template('{user_info@}{host}{:port}')
    type Authority {
        user_info: String @1
        host: String @2 @format('')
        port: Int @3
    }

    //@template("{segments:/}")
    //type Path {
    //    segments:[String] @1
    //}

    ///
    @template("{key=value:&}")
    type Query: {String: [String]}

    ///
    scheme: String @1 //@format(r'[\w\d+-.]+')
    
    ///
    authority: Authority @2
    
    ///
    path: String @3
    
    ///
    query: Query @5
    
    ///
    fragment: String @7
}

//func addQuery(url: Url, key: String, value: String)

//func addQueries(url: Url, parameters: {String: String})

//func getQuery(url: Url, key: String) -> [String]