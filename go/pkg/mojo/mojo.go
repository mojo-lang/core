package mojo

import (
    "google.golang.org/protobuf/runtime/protoimpl"
)

func GetOptionFullName(option *protoimpl.ExtensionInfo) string {
    return string(option.TypeDescriptor().FullName())
}

func EnumOptionsExtensions() []*protoimpl.ExtensionInfo {
    return []*protoimpl.ExtensionInfo{
        E_EnumAlias,
    }
}

func EnumValueOptionsExtensions() []*protoimpl.ExtensionInfo {
    return []*protoimpl.ExtensionInfo{
        E_EnumvalueAlias,
    }
}

func FileOptionsExtensions() []*protoimpl.ExtensionInfo {
    return []*protoimpl.ExtensionInfo{}
}

func MessageOptionsExtensions() []*protoimpl.ExtensionInfo {
    return []*protoimpl.ExtensionInfo{}
}

func FieldOptionsExtensions() []*protoimpl.ExtensionInfo {
    return []*protoimpl.ExtensionInfo{
        E_Alias,
        E_Key,
        E_Reference,
        E_BackReference,
        E_DbIgnore,
        E_DbJson,
        E_DbExplode,
        E_DbReference,
    }
}
