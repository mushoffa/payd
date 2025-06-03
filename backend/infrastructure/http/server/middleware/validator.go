package middleware

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	valid "github.com/go-playground/validator/v10"
)

type Validator interface {
	Struct(interface{}) error
}

type validator struct {
	i *valid.Validate
}

func NewValidator() Validator {
	_v := valid.New(valid.WithRequiredStructEnabled())

	_validator := &validator{i: _v}
	_validator.registerTagName()

	return _validator
}

func (v *validator) Struct(body interface{}) error {
	if err := v.i.Struct(body); err != nil {
		for _, err := range err.(valid.ValidationErrors) {
			return v.getErrorTag(err.Tag(), err.Field())
		}
	}
	return nil
}

func (v *validator) registerTagName() {
	v.i.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}

func (v *validator) getErrorTag(tag, field string) error {
	switch tag {
	case "required":
		return errors.New(fmt.Sprintf("Field '%s' is required.", field))

	default:
		return errors.New(fmt.Sprintf("Invalid %s data format.", field))
	}
}
