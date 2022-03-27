package core

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

const errorCodeStr = "system.408"

var errorCodeStruct = &ErrorCode{Domain: "system", Val: 408}

func TestErrorCode_Format(t *testing.T) {
    assert.Equal(t, errorCodeStr, errorCodeStruct.Format())
}

func TestErrorCode_Parse(t *testing.T) {
    ec := &ErrorCode{}
    err := ec.Parse(errorCodeStr)

    assert.NoError(t, err)
    assert.Equal(t, errorCodeStruct.Domain, ec.Domain)
    assert.Equal(t, errorCodeStruct.Val, ec.Val)
}
