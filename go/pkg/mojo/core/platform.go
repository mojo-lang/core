package core

import (
    "runtime"
    "strings"
)

const PlatformTypeName = "Platform"
const PlatformTypeFullName = "mojo.core.Platform"

func NewPlatform(os OS, arch Architecture, variant string) *Platform {
    return &Platform{
        Os:           os,
        Architecture: arch,
        Variant:      variant,
    }
}

func (m *Platform) SetOsName(name string) *Platform {
    if m != nil {
        m.OsName = name
    }
    return m
}

func (m *Platform) SetOsVersion(version string) *Platform {
    if m != nil {
        m.OsVersion = version
    }
    return m
}

func (m *Platform) IsLinux() bool {
    return m != nil && m.Os == OS_OS_LINUX
}

func (m *Platform) IsArmArch() bool {
    return m != nil && (m.Architecture == Architecture_ARCHITECTURE_ARM || m.Architecture == Architecture_ARCHITECTURE_ARM64)
}

func (m *Platform) Is64Arch() bool {
    return m != nil && (m.Architecture == Architecture_ARCHITECTURE_AMD64 || m.Architecture == Architecture_ARCHITECTURE_ARM64)
}

// https://github.com/containerd/containerd/blob/main/platforms/platforms.go
func normalizeOS(os string) string {
    if os == "" {
        return runtime.GOOS
    }
    os = strings.ToLower(os)

    switch os {
    case "macos":
        os = "darwin"
    }
    return os
}

// normalizeArch normalizes the architecture.
func normalizeArch(arch, variant string) (string, string) {
    arch, variant = strings.ToLower(arch), strings.ToLower(variant)
    switch arch {
    case "i386":
        arch = "386"
        variant = ""
    case "x86_64", "x86-64", "amd64":
        arch = "amd64"
        if variant == "v1" {
            variant = ""
        }
    case "aarch64", "arm64":
        arch = "arm64"
        switch variant {
        case "8", "v8":
            variant = ""
        }
    case "armhf":
        arch = "arm"
        variant = "v7"
    case "armel":
        arch = "arm"
        variant = "v6"
    case "arm":
        switch variant {
        case "", "7":
            variant = "v7"
        case "5", "6", "8":
            variant = "v" + variant
        }
    }

    return arch, variant
}
