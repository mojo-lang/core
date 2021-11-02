package core

import (
	"fmt"
	"github.com/google/uuid"
	jsoniter "github.com/json-iterator/go"
	"unsafe"
)

const UuidTypeName = "mojo.core.Uuid"

func init() {
	jsoniter.RegisterTypeDecoder("core.Uuid", &UuidCodec{})
	jsoniter.RegisterTypeEncoder("core.Uuid", &UuidCodec{})
}

func NewUuid() *Uuid {
	id := uuid.New()
	uuid := &Uuid{}
	copy(uuid.Value, id[:])
	return uuid
}

func ParseUuid(value string) (*Uuid, error) {
	id := &Uuid{}
	err := id.Parse(value)
	return id, err
}

func (m *Uuid) Format() string {
	return m.ToUUID().String()
}

func (m *Uuid) Parse(value string) error {
	id, err := uuid.Parse(value)
	if err == nil {
		m.Value = append(m.Value, id[:]...)
	}
	return err
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

type UuidCodec struct {
}

func (codec *UuidCodec) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	value := iter.ReadString()
	id := (*Uuid)(ptr)
	err := id.Parse(value)
	if err != nil {
		iter.ReportError("Uuid Decode", err.Error())
	}
}

func (codec *UuidCodec) IsEmpty(ptr unsafe.Pointer) bool {
	id := (*Uuid)(ptr)
	return id == nil || len(id.Value) == 0
}

func (codec *UuidCodec) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	id := (*Uuid)(ptr)
	stream.WriteVal(id.Format())
}
