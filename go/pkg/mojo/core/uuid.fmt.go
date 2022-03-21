package core

import "github.com/google/uuid"

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
        m.Val = append(m.Val, id[:]...)
    }
    return err
}
