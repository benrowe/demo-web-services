package app

import (
    "fmt"
    "github.com/go-playground/validator/v10"
    "log"
    "strings"
)

// ValidationRule
type ValidationRule string

// forPlayground converts the rule into a format for consumption by go-playground's validator
func (r ValidationRule) forPlayground() string {
    return strings.Replace(string(r), ":", "=", 0)
}

// ValidateVal the value against the rules
func ValidateVal(val interface{}, name string, rules *[]ValidationRule) (msg string, ok bool) {
    msg = ""
    ok = true

    validate := validator.New()

    s := stringifyRulesForPlayground(rules)
    if s == "" {
        log.Printf("no rules available")
        return
    }
    err := validate.Var(val, s)
    if err != nil {
        ok = false
        msg = fmt.Sprintf("The %v field is %s", name, err.Error())
    }
    return
}

// stringifyRulesForPlayground converts the the slice of validation rules into a single string for consumption with
// playgrounds validator
func stringifyRulesForPlayground(r *[]ValidationRule) string {
    var s []string
    for i, rr := range *r {
        rule := rr.forPlayground()
        if rule != "" {
            s[i] = rule
        }
    }
    return strings.Join(s, ",")
}
