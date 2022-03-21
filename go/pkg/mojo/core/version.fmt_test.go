package core

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestVersion_Parse(t *testing.T) {
    const V = "0.0.0-20211119132822-836b7b2cb4ed"
    version, err := ParseVersion(V)
    assert.NoError(t, err)
    assert.Equal(t, uint64(0), version.Major)
    assert.Equal(t, 1, len(version.PreReleases))
    assert.Equal(t, "20211119132822-836b7b2cb4ed", version.PreReleases[0])
}
