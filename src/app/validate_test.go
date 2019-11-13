package app

import (
    "github.com/stretchr/testify/assert"
    "testing"
)

func TestValidateVal(t *testing.T) {
    var msg string
    var ok bool

    // valid
    msg, ok = ValidateVal("valuessss", "value", &[]ValidationRule{"required", "min:2"})

    assert.True(t, ok)
    assert.Equal(t, "", msg)

    //first rule invalid
    msg, ok = ValidateVal("", "value", &[]ValidationRule{"required", "min:2"})

    assert.False(t, ok)
    assert.Equal(t, "The value field is required", msg)

    // second rule invalid
    msg, ok = ValidateVal("", "value", &[]ValidationRule{"required", "min:2"})

    assert.False(t, ok)
    assert.Equal(t, "The value field is required", msg)

}
