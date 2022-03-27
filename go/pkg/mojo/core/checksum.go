package core

import (
    "crypto/md5"
    "crypto/sha1"
    "crypto/sha256"
    "crypto/sha512"
    "fmt"
)

const ChecksumTypeName = "Checksum"
const ChecksumTypeFullName = "mojo.core.Checksum"

var checksumLength = map[Checksum_Algorithm]int{
    Checksum_ALGORITHM_MD5:    32,
    Checksum_ALGORITHM_SHA1:   40,
    Checksum_ALGORITHM_SHA256: 64,
    Checksum_ALGORITHM_SHA512: 128,
}

func NewChecksum(algorithm Checksum_Algorithm, data []byte) *Checksum {
    return (&Checksum{Algorithm: algorithm}).Sum(data)
}

func (m *Checksum) Check(data []byte) bool {
    if m != nil && m.Algorithm > 0 && len(m.Value) > 0 {
        switch m.Algorithm {
        case Checksum_ALGORITHM_MD5:
            return m.Value == fmt.Sprintf("%x", md5.Sum(data))
        case Checksum_ALGORITHM_SHA1:
            return m.Value == fmt.Sprintf("%x", sha1.Sum(data))
        case Checksum_ALGORITHM_SHA256:
            return m.Value == fmt.Sprintf("%x", sha256.Sum256(data))
        case Checksum_ALGORITHM_SHA512:
            return m.Value == fmt.Sprintf("%x", sha512.Sum512(data))
        }
    }
    return false
}

func (m *Checksum) Sum(data []byte) *Checksum {
    if m != nil {
        if m.Algorithm == 0 {
            m.Algorithm = Checksum_ALGORITHM_SHA256
        }
        switch m.Algorithm {
        case Checksum_ALGORITHM_MD5:
            m.Value = fmt.Sprintf("%x", md5.Sum(data))
        case Checksum_ALGORITHM_SHA1:
            m.Value = fmt.Sprintf("%x", sha1.Sum(data))
        case Checksum_ALGORITHM_SHA256:
            m.Value = fmt.Sprintf("%x", sha256.Sum256(data))
        case Checksum_ALGORITHM_SHA512:
            m.Value = fmt.Sprintf("%x", sha512.Sum512(data))
        }
    }
    return m
}

func (m *Checksum) IsValid() bool {
    if !m.IsEmpty() {
        return len(m.Value) == checksumLength[m.Algorithm]
    }
    return false
}

func (m *Checksum) IsEmpty() bool {
    return m == nil || m.Algorithm == 0 || len(m.Value) == 0
}
