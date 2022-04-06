package core

import (
    "github.com/stretchr/testify/assert"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "os"
    "path/filepath"
    "testing"
)

type User2 struct {
    ID     uint
    Name   string
    Values *Values
}

func TestValues_Scan(t *testing.T) {
    DB, err := gorm.Open(sqlite.Open(filepath.Join(os.TempDir(), "gorm.db")), &gorm.Config{})
    DB = DB.Debug()
    if err != nil {
        t.Errorf("failed to connect database")
    }

    user := User2{Name: "mojo", Values: NewValues().AppendString("foo").AppendInt64(100).AppendString("bar")}
    DB.Migrator().DropTable(&User2{})
    DB.AutoMigrate(&User2{})
    DB.Save(&user)

    var count int64

    if DB.Model(&User2{}).Where("name = ?", user.Name).Count(&count).Error != nil || count != 1 {
        t.Errorf("Count soft deleted record, expects: %v, got: %v", 1, count)
    }

    var values Values
    err = DB.Model(&User2{}).Select("values").Where("name = ?", user.Name).Scan(&values).Error
    assert.NoError(t, err)

    assert.Equal(t, 3, values.Len())
    assert.Equal(t, "foo", values.GetString(0))
    assert.Equal(t, int64(100), values.GetInt64(1))
}
