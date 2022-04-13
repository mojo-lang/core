package core

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

var platforms = map[string]*Platform{
    "linux/arm64": {
        Os:           OS_OS_LINUX,
        Architecture: Architecture_ARCHITECTURE_ARM64,
    },
}

func TestPlatformStringCodec_Decode(t *testing.T) {
    for k, p := range platforms {
        platform, err := ParsePlatform(k)
        assert.NoError(t, err)
        assert.Equal(t, p.Os, platform.Os)
        assert.Equal(t, p.Architecture, platform.Architecture)
    }
}
