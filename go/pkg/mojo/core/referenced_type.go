package core

type ReferencedType interface {
	FieldMask() map[string]bool

	// return the referenced value
	Referenced() interface{}
}
