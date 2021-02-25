
type Image {
    title: String @1
    url: Url @2
    type: String @3
    properties: {String: Value} @4

    timestamp: Timestamp @10
}
