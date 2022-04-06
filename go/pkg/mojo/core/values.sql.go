package core

import (
    "database/sql/driver"
    "fmt"
    jsoniter "github.com/json-iterator/go"

    "gorm.io/gorm"
    "gorm.io/gorm/schema"
)

func (x *Values) Value() (driver.Value, error) {
    if x != nil {
        return jsoniter.ConfigFastest.MarshalToString(x)
    }
    return nil, nil
}

func (x *Values) Scan(src interface{}) error {
    if x != nil && src != nil {
        switch v := src.(type) {
        case []byte:
            return jsoniter.ConfigFastest.Unmarshal(v, x)
        case string:
            return jsoniter.ConfigFastest.UnmarshalFromString(v, x)
        default:
            return fmt.Errorf("failed to unmarshal JSON src: %v", src)
        }
    }
    return nil
}

func (x *Values) GormDBDataType(db *gorm.DB, field *schema.Field) string {
    return GormDBJSONType(db)
}
