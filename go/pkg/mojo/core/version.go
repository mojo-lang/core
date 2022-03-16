package core

import (
    "regexp"
)

const VersionTypeName = "Version"
const VersionTypeFullName = "mojo.core.Version"

func NewVersion(major int, minor int, patch int) *Version {
    return &Version{
        Major: uint64(major),
        Minor: uint64(minor),
        Patch: uint64(patch),
    }
}

// IsVersionTag
// v1
// v1alpha
// v1beta
func IsVersionTag(version string) bool {
    matched, _ := regexp.MatchString(`v[0-9]+[a-zA-z]*[0-9]*`, version)
    return matched
}
