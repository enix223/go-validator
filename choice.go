package validators

import (
	"errors"
	"reflect"
)

const (
	errChoicesNotSlice = "choices should be with type Array, Slice"
)

// ChoiceValidator validator to validate the value in given choice values
type ChoiceValidator struct {
	ValidatorName string
	Value         interface{}
	Message       string
	Choices       interface{}
}

func validateChoices(choices interface{}) {
	val := reflect.ValueOf(choices)
	k := val.Kind()
	if k != reflect.Slice && k != reflect.Array {
		panic(errChoicesNotSlice)
	}
}

// NewChoiceValidator create a blank validator
// choices should be with type Array, Slice or string or it will panic
func NewChoiceValidator(name, errorMessage string, validateValue, choices interface{}) *ChoiceValidator {
	validateChoices(choices)
	return &ChoiceValidator{
		ValidatorName: name,
		Message:       errorMessage,
		Value:         validateValue,
		Choices:       choices,
	}
}

// Validate validate `value` is blank or not
func (v *ChoiceValidator) Validate(value interface{}) error {
	validateChoices(v.Choices)
	val := reflect.ValueOf(v.Choices)
	valid := false
	for i := 0; i < val.Len(); i++ {
		c := val.Index(i)
		if reflect.DeepEqual(value, c.Interface()) {
			valid = true
			break
		}
	}

	if !valid {
		return errors.New(v.Message)
	}

	return nil
}

// Identifier get ValidateHandler identifier
func (v *ChoiceValidator) Identifier() ValidateIdentifier {
	return ValidateIdentifier(v.ValidatorName)
}

// SetValue set value to be validated
func (v *ChoiceValidator) SetValue(value interface{}) {
	v.Value = value
}

// GetValue get the value to be validated
func (v *ChoiceValidator) GetValue() interface{} {
	return v.Value
}
