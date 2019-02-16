package validators

import (
	"errors"
)

// ValidateFunc function to validate the input value,
// return true if valid, otherwise return false
type ValidateFunc func(value interface{}) bool

// FunctionValidator validate value with given function
type FunctionValidator struct {
	ValidatorName string
	Value         interface{}
	Message       string
	Func          ValidateFunc
}

// NewValidateFunc create a function validator
func NewValidateFunc(name, errorMessage string, validateValue interface{}, function ValidateFunc) *FunctionValidator {
	return &FunctionValidator{
		ValidatorName: name,
		Message:       errorMessage,
		Value:         validateValue,
		Func:          function,
	}
}

// Validate validate `value` is blank or not
func (v *FunctionValidator) Validate(value interface{}) error {
	if v.Func(value) {
		return nil
	}

	return errors.New(v.Message)
}

// Identifier get ValidateHandler identifier
func (v *FunctionValidator) Identifier() ValidateIdentifier {
	return ValidateIdentifier(v.ValidatorName)
}

// SetValue set value to be validated
func (v *FunctionValidator) SetValue(value interface{}) {
	v.Value = value
}

// GetValue get the value to be validated
func (v *FunctionValidator) GetValue() interface{} {
	return v.Value
}
