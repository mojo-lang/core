
@format('{vals:','}')
type Ordering {
    enum Sort {
        unspecified @0
        asc  @1
        desc @2
    }

    @format('{field}{ sort}')
    type Value {
        field: String @1
        sort: Sort @2
    }

    vals: [Value] @1
}
