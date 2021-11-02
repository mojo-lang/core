package core

import (
	"strconv"
)

func NewId(value interface{}) *Id {
	switch v := value.(type) {
	case uint64:
		return NewIntId(v)
	case string:
		return NewStringId(v)
	case *Uuid:
		return NewUuidId(v)
	}
	return nil
}

func NewIntId(value uint64) *Id {
	return &Id{Id: &Id_Uint64Val{Uint64Val: value}}
}

func NewStringId(value string) *Id {
	return &Id{Id: &Id_StringVal{StringVal: value}}
}

func NewUuidId(value *Uuid) *Id {
	return &Id{Id: &Id_UuidVal{UuidVal: value}}
}

func UuidId() *Id {
	return NewUuidId(NewUuid())
}

func ParseId(id string) (*Id, error) {
	i := &Id{}
	err := i.Parse(id)
	if err != nil {
		return nil, err
	}
	return i, nil
}

func (m *Id) SetInt(id uint64) {
	m.Id = &Id_Uint64Val{Uint64Val: id}
}

func (m *Id) SetString(id string) {
	m.Id = &Id_StringVal{StringVal: id}
}

func (m *Id) SetUuid(id *Uuid) {
	m.Id = &Id_UuidVal{UuidVal: id}
}

func (m *Id) SetUuidString(id string) {
	uuid, err := ParseUuid(id)
	if err == nil {
		m.SetUuid(uuid)
	}
}

func (m *Id) RegenerateUuid() {
	m.SetUuid(NewUuid())
}

func (m *Id) Format() string {
	switch x := m.Id.(type) {
	case *Id_Uint64Val:
		return strconv.FormatUint(x.Uint64Val, 10)
	case *Id_StringVal:
		return x.StringVal
	case *Id_UuidVal:
		return x.UuidVal.Format()
	default:
		return ""
	}
}

func (m *Id) Parse(value string) error {
	v, err := strconv.Atoi(value)
	if err == nil {
		m.SetInt(uint64(v))
	} else {
		uuid, err := ParseUuid(value)
		if err != nil {
			m.SetUuid(uuid)
		} else {
			m.SetString(value)
		}
	}

	return err
}
