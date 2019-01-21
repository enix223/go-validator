package validators

import (
	"errors"
	"reflect"
)

// BlankValidator validator for blank value
type BlankValidator struct {
	ValidatorName string
	Value         interface{}
	Message       string
	BlankValue    interface{}
}

// NewBlankValidator create a blank validator
func NewBlankValidator(name, errorMessage string, validateValue, blankValue interface{}) *BlankValidator {
	return &BlankValidator{
		ValidatorName: name,
		Message:       errorMessage,
		Value:         validateValue,
		BlankValue:    blankValue,
	}
}

// Validate validate `value` is blank or not
func (v *BlankValidator) Validate(value interface{}) error {
	if reflect.DeepEqual(value, v.BlankValue) {
		return errors.New(v.Message)
	}

	return nil
}

// Identifier get ValidateHandler identifier
func (v *BlankValidator) Identifier() ValidateIdentifier {
	return ValidateIdentifier(v.ValidatorName)
}

// SetValue set value to be validated
func (v *BlankValidator) SetValue(value interface{}) {
	v.Value = value
}

// GetValue get the value to be validated
func (v *BlankValidator) GetValue() interface{} {
	return v.Value
}
