package core

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
