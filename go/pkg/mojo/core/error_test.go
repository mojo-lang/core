package core

import (
    jsoniter "github.com/json-iterator/go"
    "github.com/stretchr/testify/assert"
    "google.golang.org/protobuf/proto"
    "testing"
)

func TestIsError(t *testing.T) {
    assert.True(t, IsError(NewErrorFrom(400, "msg")))
}

func TestError_AddDetail(t *testing.T) {
    err := NewNotFoundError("not found")
    err.AddDetail(&RetryInfo{RetryDelay: NewDurationFromSeconds(100)})

    str, e := jsoniter.Marshal(err)
    assert.NoError(t, e)

    notFoundErr := &Error{}
    e = jsoniter.Unmarshal(str, notFoundErr)
    assert.NoError(t, e)

    pb, e := proto.Marshal(err)
    assert.NoError(t, e)

    pb2, e := proto.Marshal(err.ToError())
    assert.NoError(t, e)
    assert.Equal(t, pb, pb2)
}

func TestError_AddDetail2(t *testing.T) {
    err := NewNotFoundError("not found")
    err.AddDetail(NewAny(&RetryInfo{RetryDelay: NewDurationFromSeconds(100)}))

    assert.True(t, err.ContainsDetailType(&RetryInfo{}))
}

func TestError_ContainsDetailType(t *testing.T) {
    err := NewNotFoundError("not found")
    err.AddDetail(&RetryInfo{RetryDelay: NewDurationFromSeconds(100)})

    assert.True(t, err.ContainsDetailType(&RetryInfo{}))
    assert.False(t, err.ContainsDetailType(&ErrorInfo{}))
}
