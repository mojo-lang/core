package core

import "regexp"

func ParseRegex(fieldMask string) (*FieldMask, error) {
	mask := NewFieldMask(fieldMask)
	return mask, nil
}

func (x *Regex) Format() string {
	if x != nil {
		return x.Expression
	} else {
		return ""
	}
}

func (x *Regex) ToString() string {
	return x.Format()
}

func (x *Regex) Parse(expr string) error {
	if x != nil {
		if _, err := regexp.Compile(expr); err != nil {
			return err
		}

		x.Expression = expr
	}
	return nil
}
