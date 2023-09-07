package core

import (
	"regexp"

	"github.com/mojo-lang/core/go/pkg/logs"
)

const RegexTypeName = "Regex"
const RegexTypeFullName = "mojo.core.Regex"

func NewRegex(expr string) *Regex {
	if _, err := regexp.Compile(expr); err == nil {
		return &Regex{Expression: expr}
	}

	return nil
}

// ToRegexp convert the Regex to *regexp.Regexp
func (x *Regex) ToRegexp() *regexp.Regexp {
	if x != nil {
		if len(x.Expression) > 0 {
			r, err := regexp.Compile(x.Expression)
			if err == nil {
				return r
			} else {
				logs.Warnw("failed to compile the regex", "regex", x.Expression, "error", err)
			}
		}
	}
	return nil
}
