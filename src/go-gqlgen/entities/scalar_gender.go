package entities

import (
	"fmt"
	"io"
	"strconv"
)

type Gender string

const (
	GenderMale      Gender = "MALE"
	GenderFemale    Gender = "FEMALE"
	GenderUndefined Gender = "UNDEFINED"
)

var AllGender = []Gender{
	GenderMale,
	GenderFemale,
	GenderUndefined,
}

func (e Gender) IsValid() bool {
	switch e {
	case GenderMale, GenderFemale, GenderUndefined:
		return true
	}
	return false
}

func (e Gender) String() string {
	return string(e)
}

func (e *Gender) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Gender(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Gender", str)
	}
	return nil
}

func (e Gender) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
