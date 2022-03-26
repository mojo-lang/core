package core

import "os"

// CreateDir create dir with recursively
func CreateDir(filePath string) error {
    if !IsExist(filePath) {
        return os.MkdirAll(filePath, os.ModePerm)
    }
    return nil
}

// IsExist check the file or directory exist, return true if exist
func IsExist(path string) bool {
    _, err := os.Stat(path) //os.Stat get file info
    if err != nil {
        if os.IsExist(err) {
            return true
        }
        return false
    }
    return true
}
