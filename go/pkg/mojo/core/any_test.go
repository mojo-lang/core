package core

import (
    "github.com/stretchr/testify/assert"
    "google.golang.org/protobuf/proto"
    "testing"
)

var anyScalarValues = []interface{}{
    int64(100),
    "this is a string",
}

func TestAny_Marshal_Scalar(t *testing.T) {
    for _, v := range anyScalarValues {
        any := NewAny(v)
        b, err := proto.Marshal(any)
        assert.NoError(t, err)

        a := &Any{}
        err = proto.Unmarshal(b, a)
        assert.NoError(t, err)

        assert.Equal(t, v, a.Get())
    }
}

func TestAny_Marshal_Struct(t *testing.T) {
    any := NewAny(&Checksum{
        Algorithm: Checksum_ALGORITHM_MD5,
        Val:       "md5",
    })

    b, err := proto.Marshal(any)
    assert.NoError(t, err)

    a := &Any{}
    err = proto.Unmarshal(b, a)
    assert.NoError(t, err)

    v, ok := a.Get().(*Checksum)
    assert.True(t, ok)
    assert.Equal(t, Checksum_ALGORITHM_MD5, v.Algorithm)
    assert.Equal(t, "md5", v.Val)
}
