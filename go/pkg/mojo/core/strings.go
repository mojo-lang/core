package core

import "strings"

func RemoveDoubleQuote(str string) string {
	return RemoveQuote(str, `"`)
}

func RemoveSingleQuote(str string) string {
	return RemoveQuote(str, `'`)
}

func RemoveQuote(str string, quote string) string {
	if strings.HasPrefix(str, quote) && strings.HasSuffix(str, quote) {
		return strings.TrimSuffix(strings.TrimPrefix(str, quote), quote)
	}
	return str
}
