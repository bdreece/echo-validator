package echovalidator

import (
	"github.com/go-playground/validator/v10"
)

type Validator struct {
    validator *validator.Validate
}

func (v *Validator) Validate(i any) error {
    return v.validator.Struct(i)
}

func New() *Validator {
    return &Validator{validator.New()}
}

var Default = New()
