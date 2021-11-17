package strcase

import "github.com/iancoleman/strcase"

// ToCamel converts a string to CamelCase
func ToCamel(s string) string {
	return strcase.ToCamel(s)
}

// ToLowerCamel converts a string to lowerCamelCase
func ToLowerCamel(s string) string {
	return strcase.ToLowerCamel(s)
}

// ConfigureAcronym allows you to add additional words which will be considered acronyms
func ConfigureAcronym(key, val string) {
	strcase.ConfigureAcronym(key, val)
}
