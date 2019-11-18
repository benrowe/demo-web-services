package app

import (
    "github.com/go-playground/locales"
    locale_en "github.com/go-playground/locales/en"
    ut "github.com/go-playground/universal-translator"
    "github.com/go-playground/validator/v10"
    en_translations "github.com/go-playground/validator/v10/translations/en"
    "log"
    "strings"
)

var (
    validate *validator.Validate
    trans    ut.Translator
)

// get validator instance
func get() *validator.Validate {
    if validate == nil {
        validate = newInstance("en")
    }
    return validate
}

// newInstance validator instance
func newInstance(l string) *validator.Validate {
    availLocales := []locales.Translator{locale_en.New()}
    uni := ut.New(availLocales[0], availLocales...)
    t, found := uni.GetTranslator("en")
    if found == false {
        log.Printf("unable to get translator %s", l)
    }
    trans = t

    v := validator.New()
    err := en_translations.RegisterDefaultTranslations(v, trans)
    if err != nil {
        log.Printf("crap %v", err)
    }
    err = v.RegisterTranslation("required", t, func(ut ut.Translator) error {
        return ut.Add("required", "The {0} field is required", true)
    }, func(ut ut.Translator, fe validator.FieldError) string {
        t, _ := ut.T("required", fe.Field())

        return t
    })

    if err != nil {
        log.Printf("unable to register translations for validator: %v", err)
    }

    return v
}

// ValidationRule
type ValidationRule string

// forPlayground converts the rule into a format for consumption by go-playground's validator
func (r ValidationRule) forPlayground() string {
    return strings.Replace(string(r), ":", "=", -1)
}

// ValidateVal the value against the rules
func ValidateVal(val interface{}, name string, rules *[]ValidationRule) (msg string, ok bool) {
    msg = ""
    ok = true

    validate := get()

    s := stringifyRulesForPlayground(rules)
    if s == "" {
        log.Printf("no rules available")
        return
    }
    err := validate.Var(val, s)
    if err != nil {
        ok = false
        errs := err.(validator.ValidationErrors)
        for _, e := range errs {
            // can translate each error one at a time.
            msg = strings.Replace(e.Translate(trans), "  ", " "+name+" ", 1)
        }

        //msg = errs.Error()
    }
    return
}

// stringifyRulesForPlayground converts the the slice of validation rules into a single string for consumption with
// playgrounds validator
func stringifyRulesForPlayground(r *[]ValidationRule) string {
    var s []string
    //s := []string{}
    for _, rr := range *r {
        rule := rr.forPlayground()
        if rule != "" {
            s = append(s, rule)
        }
    }
    return strings.Join(s, ",")
}
