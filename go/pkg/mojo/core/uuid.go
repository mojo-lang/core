package core

import (
	"fmt"
	"github.com/google/uuid"
)

const UuidTypeName = "mojo.core.Uuid"

func NewUuid() *Uuid {
	id := uuid.New()
	uuid := &Uuid{}
	copy(uuid.Value, id[:])
	return uuid
}

func (m *Uuid) ToUUID() uuid.UUID {
	if m != nil && len(m.Value) > 0 {
		id, err := uuid.ParseBytes(m.Value)
		if err != nil {
			return uuid.New()
		}
		return id
	}
	return uuid.New()
}

//// MarshalText implements encoding.TextMarshaler.
//func (m *Uuid) MarshalText() ([]byte, error) {
//	var js [36]byte
//	encodeHex(js[:], uuid)
//	return js[:], nil
//}
//
//// UnmarshalText implements encoding.TextUnmarshaler.
//func (m *Uuid UnmarshalText(data []byte) error {
//	id, err := ParseBytes(data)
//	if err != nil {
//		return err
//	}
//	*uuid = id
//	return nil
//}

// MarshalBinary implements encoding.BinaryMarshaler.
func (m *Uuid) MarshalBinary() ([]byte, error) {
	return m.Value[:], nil
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler.
func (m *Uuid) UnmarshalBinary(data []byte) error {
	if len(data) != 16 {
		return fmt.Errorf("invalid UUID (got %d bytes)", len(data))
	}
	copy(m.Value[:], data)
	return nil
}


