package core

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestCaseStyle_Format(t *testing.T) {
    var cs CaseStyle
    assert.Empty(t, cs.Format())
}
