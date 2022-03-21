package core

type StringType interface {
    AsStringType() string
}

type ScalarType interface {
    AsScalarType() interface{}
}

type ArrayType interface {
    AsArrayType() interface{}
}

type MapType interface {
    AsMapType() interface{}
}

type CollectionType interface {
    AsCollectionType() interface{}
}
