package validators

import (
	"testing"
)

const assertTemplate = "expected value: %v, got: %v"

func TestBlankStringValidator(t *testing.T) {
	validator := &BlankValidator{
		ValidatorName: "abc",
		Value:         "123",
		Message:       "error",
		BlankValue:    "",
	}

	// should not raise blank error
	err := validator.Validate(validator.Value)
	if err != nil {
		t.Fatalf(assertTemplate, nil, err)
	}

	validator = &BlankValidator{
		ValidatorName: "abc",
		Value:         "",
		Message:       "blank_error",
		BlankValue:    "",
	}

	// should raise blank error
	err = validator.Validate(validator.Value)
	if err == nil || err.Error() != "blank_error" {
		t.Fatalf(assertTemplate, "blank_error", nil)
	}
}

func TestBlankIntValidator(t *testing.T) {
	validator := &BlankValidator{
		ValidatorName: "abc",
		Value:         1,
		Message:       "error",
		BlankValue:    0,
	}

	err := validator.Validate(validator.Value)
	if err != nil {
		t.Fatalf(assertTemplate, nil, err)
	}

	validator = &BlankValidator{
		ValidatorName: "abc",
		Value:         0,
		Message:       "blank_error",
		BlankValue:    0,
	}

	err = validator.Validate(validator.Value)
	if err == nil || err.Error() != "blank_error" {
		t.Fatalf(assertTemplate, "blank_error", err)
	}
}

func TestBlankFloatValidator(t *testing.T) {
	validator := &BlankValidator{
		ValidatorName: "abc",
		Value:         1.0,
		Message:       "error",
		BlankValue:    0.0,
	}

	err := validator.Validate(validator.Value)
	if err != nil {
		t.Fatalf(assertTemplate, nil, err)
	}

	validator = &BlankValidator{
		ValidatorName: "abc",
		Value:         0.0,
		Message:       "blank_error",
		BlankValue:    0.0,
	}

	err = validator.Validate(validator.Value)
	if err == nil || err.Error() != "blank_error" {
		t.Fatalf(assertTemplate, "blank_error", err)
	}
}

func TestBoolValidator(t *testing.T) {
	validator := &BlankValidator{
		ValidatorName: "abc",
		Value:         false,
		Message:       "error",
		BlankValue:    true,
	}

	err := validator.Validate(validator.Value)
	if err != nil {
		t.Fatalf(assertTemplate, nil, err)
	}

	validator = &BlankValidator{
		ValidatorName: "abc",
		Value:         false,
		Message:       "blank_error",
		BlankValue:    false,
	}

	err = validator.Validate(validator.Value)
	if err == nil || err.Error() != "blank_error" {
		t.Fatalf(assertTemplate, "blank_error", err)
	}
}

func TestNewBlankValidator(t *testing.T) {
	validator := NewBlankValidator("abc", "err", "1", "2")
	if validator.ValidatorName != "abc" {
		t.Fatalf(assertTemplate, "abc", validator.ValidatorName)
	}

	if validator.Message != "err" {
		t.Fatalf(assertTemplate, "err", validator.Message)
	}

	if validator.Value != "1" {
		t.Fatalf(assertTemplate, "1", validator.Value)
	}

	if validator.BlankValue != "2" {
		t.Fatalf(assertTemplate, "2", validator.BlankValue)
	}
}

func TestIdentifier(t *testing.T) {
	validator := &BlankValidator{
		ValidatorName: "abc",
		Value:         "1",
		Message:       "error",
		BlankValue:    "2",
	}
	if validator.Identifier() != "abc" {
		t.Fatalf(assertTemplate, "abc", validator.ValidatorName)
	}
}

func TestSetValue(t *testing.T) {
	validator := &BlankValidator{
		ValidatorName: "abc",
		Value:         "1",
		Message:       "error",
		BlankValue:    "2",
	}
	validator.SetValue("3")
	if validator.Value != "3" {
		t.Fatalf(assertTemplate, "3", validator.Value)
	}
}

func TestGetValue(t *testing.T) {
	validator := &BlankValidator{
		ValidatorName: "abc",
		Value:         "1",
		Message:       "error",
		BlankValue:    "2",
	}
	if validator.GetValue() != "1" {
		t.Fatalf(assertTemplate, "1", validator.Value)
	}
}
