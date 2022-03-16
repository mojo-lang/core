package core

import "reflect"

const UnionTypeName = "Union"
const UnionTypeFullName = "mojo.core.Union"

type IsUnion interface {
    IsUnion()
}

func GetUnionPrimeType(union interface{}) interface{} {
    if union != nil {
        value := reflect.ValueOf(union)
        for {
            if _, ok := value.Interface().(IsUnion); ok {
                value = reflect.Indirect(value).Field(0).Elem()
                value = reflect.Indirect(value).Field(0)
            } else {
                break
            }
        }
        return value.Interface()
    }
    return nil
}

func GetUnionField(union interface{}, field string) reflect.Value {
    if union != nil {
        value := reflect.ValueOf(union)
        for {
            if _, ok := value.Interface().(IsUnion); ok {
                value = reflect.Indirect(value).Field(0).Elem()
                value = reflect.Indirect(value).Field(0)
            } else {
                break
            }
        }
        return reflect.Indirect(value).FieldByName(field)
    }
    return reflect.ValueOf(nil)
}

func SetUnionField(union interface{}, field string, val reflect.Value) {
    if union != nil {
        value := reflect.ValueOf(union)
        for {
            if _, ok := value.Interface().(IsUnion); ok {
                value = reflect.Indirect(value).Field(0).Elem()
                value = reflect.Indirect(value).Field(0)
            } else {
                break
            }
        }

        value = reflect.Indirect(value).FieldByName(field)
        if value.CanSet() {
            value.Set(val)
        }
    }
}
