///
/// `Any` contains an arbitrary serialized message along with a URL that describes the type of the serialized message.
///
type Any {
    type:    Url @1 @alias("@type")

    /// Must be valid serialized data of the above specified type.
    value: Bytes @2
}
