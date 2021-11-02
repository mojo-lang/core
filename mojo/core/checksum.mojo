
/// Checksum
/// 
/// MD5, SHA1, SHA256 and SHA512.
@format('{algorithm} {value}')
type Checksum {
    /// commonly used algorithm for Checksum
    enum Algorithm {
        unspecified @0

        md5 @1    //<
        sha1 @2   //<
        sha256 @3 //<
        sha512 @4 //<
    }

    /// the algorithm of the checksum
    algorithm: Algorithm @1
    
    /// the checksum value
    value: String @2
}
