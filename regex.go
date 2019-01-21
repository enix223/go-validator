package validators

import (
	"errors"
	"reflect"
	"regexp"
)

// RegExpValidator regular expression validator
type RegExpValidator struct {
	ValidatorName string
	Value         interface{}
	Message       string
	Pattern       *regexp.Regexp
}

// NewRegExpValidator create a regexp validator
func NewRegExpValidator(
	name,
	errorMessage string,
	validateValue interface{},
	pattern *regexp.Regexp) *RegExpValidator {
	return &RegExpValidator{
		ValidatorName: name,
		Message:       errorMessage,
		Value:         validateValue,
		Pattern:       pattern,
	}
}

// Validate validate the input value, if valid, then return nil as error,
// otherwise, return the error for the validation failure reason
func (v *RegExpValidator) Validate(value interface{}) error {
	val := reflect.ValueOf(value)
	if val.Kind() != reflect.String {
		// not a valid string
		return errors.New(v.Message)
	}

	valStr := val.String()
	if !v.Pattern.MatchString(valStr) {
		// not match
		return errors.New(v.Message)
	}

	return nil
}

// Identifier get ValidateHandler identifier
func (v *RegExpValidator) Identifier() ValidateIdentifier {
	return ValidateIdentifier(v.ValidatorName)
}

// SetValue set value to be validated
func (v *RegExpValidator) SetValue(value interface{}) {
	v.Value = value
}

// GetValue get the value to be validated
func (v *RegExpValidator) GetValue() interface{} {
	return v.Value
}
