package app

import (
    "github.com/thedevsaddam/govalidator"
    "log"
    "strings"
)

type validateValue struct {
    val interface{}
}

// Validate the value against the rules
func Validate(val interface{}, name string, rules []string) (msg string, ok bool) {
    msg = ""
    ok = true

    r := govalidator.MapData{
        "val": rules,
    }
    opt := govalidator.Options{
        Data:  &validateValue{val: val},
        Rules: r,
    }

    v := govalidator.New(opt)
    e := v.ValidateStruct()

    if len(e) > 0 {
        log.Printf("%+v", e)
        // we have an error
        msg = strings.Replace(e["val"][0], "val", name, 1)
        ok = false
    }

    return
}
