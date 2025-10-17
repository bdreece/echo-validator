// Package echovalidator provides a wrapper for [github.com/goplayground/validator/v10.Validate]
// implementing [github.com/labstack/echo/v4.Validator].
package echovalidator

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// The [validator.Validate] wrapper.
type Validator struct {
	validator *validator.Validate
}

// Validate implements [echo.Validator]
func (v *Validator) Validate(i any) error {
	return v.validator.Struct(i)
}

// Creates a new [Validator].
func New(v *validator.Validate) *Validator {
	return &Validator{v}
}

// A global singleton [Validator].
var Default = New(validator.New())

var _ echo.Validator = (*Validator)(nil)
