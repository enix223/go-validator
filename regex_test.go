package validators

import (
	"regexp"
	"testing"
)

func TestRegexpValidator(t *testing.T) {
	validator := &RegExpValidator{
		ValidatorName: "abc",
		Value:         "1234",
		Message:       "err",
		Pattern:       regexp.MustCompile("^\\d+$"),
	}

	err := validator.Validate(validator.Value)
	if err != nil {
		t.Fatalf(assertTemplate, nil, err)
	}

	err = validator.Validate("abc")
	if err == nil {
		t.Fatalf(assertTemplate, "err", nil)
	}

	err = validator.Validate(123)
	if err == nil {
		t.Fatalf(assertTemplate, "err", nil)
	}

	err = validator.Validate(true)
	if err == nil {
		t.Fatalf(assertTemplate, "err", nil)
	}
}

func TestRegexIdentifier(t *testing.T) {
	validator := &RegExpValidator{
		ValidatorName: "abc",
	}
	if validator.Identifier() != "abc" {
		t.Fatalf(assertTemplate, "abc", validator.ValidatorName)
	}
}

func TestRegexSetValue(t *testing.T) {
	validator := &RegExpValidator{}
	validator.SetValue("3")
	if validator.Value != "3" {
		t.Fatalf(assertTemplate, "3", validator.Value)
	}
}

func TestRegexGetValue(t *testing.T) {
	validator := &RegExpValidator{
		Value: "1",
	}
	if validator.GetValue() != "1" {
		t.Fatalf(assertTemplate, "1", validator.Value)
	}
}
