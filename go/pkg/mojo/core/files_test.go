package core

import (
    "github.com/stretchr/testify/assert"
    "os"
    "path"
    "syscall"
    "testing"
)

func TestMkDir(t *testing.T) {
    dir := path.Join(os.TempDir(), "dir")

    err := MkDir(dir, 0777)
    assert.NoError(t, err)

    defer func() {
        os.Remove(dir)
    }()

    info, err := os.Stat(dir)
    assert.NoError(t, err)
    assert.Equal(t, os.FileMode(0777), info.Mode()&os.ModePerm)
}

func TestCreateDir(t *testing.T) {
    dir := path.Join(os.TempDir(), "dir")

    err := CreateDir(dir)
    assert.NoError(t, err)

    defer func() {
        os.Remove(dir)
    }()

    mask := syscall.Umask(0)
    syscall.Umask(mask)

    info, err := os.Stat(dir)
    assert.NoError(t, err)
    assert.Equal(t, os.FileMode(0777^mask), info.Mode()&os.ModePerm)
}
