package core

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestParseMonth(t *testing.T) {
    mon, err := ParseMonth("May")
    assert.NoError(t, err)
    assert.Equal(t, Month_MONTH_MAY, mon)

    mon, err = ParseMonth("may")
    assert.NoError(t, err)
    assert.Equal(t, Month_MONTH_MAY, mon)
}
