package core

import "github.com/mojo-lang/core/go/pkg/mojo/core/strcase"

const IrregularCaseRuleAttributeName = "irregular_case_rule"

func ApplyIrregularCaseRuleAttribute(rule *Regex, replacement string) error {
    if rule != nil && len(rule.Expression) > 0 && len(replacement) > 0 {
        strcase.RegisterIrregularCaseRule(rule.Expression, replacement)
    }
    return nil
}
