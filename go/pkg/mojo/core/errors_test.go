package core

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAbortedError(t *testing.T) {
	err := NewAbortedError("test %s", "value")
	assert.Equal(t, Aborted.Code, err.Code.Code)
	assert.Equal(t, "test value", err.Message)
}

func TestNewAbortedError2(t *testing.T) {
	err := NewAbortedError("test value")
	assert.Equal(t, Aborted.Code, err.Code.Code)
	assert.Equal(t, "test value", err.Message)
}

func TestIsAbortedError(t *testing.T) {
	err := NewAbortedError("aborted")
	assert.True(t, IsAbortedError(err))
}

func TestIsCancelledError(t *testing.T) {
	assert.True(t, IsCancelledError(NewCancelledError("cancel")))
}

func TestNewBadRequestError(t *testing.T) {
	err := NewBadRequestError("bad request")
	assert.Equal(t, BadRequest, err.Code)
}

func TestErr_AddDetail(t *testing.T) {
	err := NewNotFoundError("not found")
	err.AddDetail(&RetryInfo{RetryDelay: NewDurationFromSeconds(100)})

	str, e := jsoniter.Marshal(err)
	assert.NoError(t, e)

	notFoundErr := &Error{}
	e = jsoniter.Unmarshal(str, notFoundErr)
	assert.NoError(t, e)
}
