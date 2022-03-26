package core

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

const platformStr = "linux/amd64-Ubuntu/14.02"

var platformStruct = &Platform{
    Architecture: Architecture_ARCHITECTURE_AMD64,
    Variant:      "",
    Os:           OS_OS_LINUX,
    OsName:       "Ubuntu",
    OsVersion:    "14.02",
}

func TestPlatform_Format(t *testing.T) {
    assert.Equal(t, platformStr, platformStruct.Format())
}

func TestPlatform_Parse(t *testing.T) {
    platform := &Platform{}
    err := platform.Parse(platformStr)
    assert.NoError(t, err)
    assert.Equal(t, platformStruct.Architecture, platform.Architecture)
    assert.Equal(t, platformStruct.Os, platform.Os)
    assert.Equal(t, platformStruct.OsName, platform.OsName)
    assert.Equal(t, platformStruct.OsVersion, platform.OsVersion)
}
