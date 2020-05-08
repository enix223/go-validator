package validators

import (
	"reflect"
	"testing"
)

func TestChoiceValidator(t *testing.T) {
	vs := []*ChoiceValidator{
		{
			ValidatorName: "choices",
			Value:         1,
			Message:       "should be in [1, 2, 3]",
			Choices:       [3]int{1, 2, 3},
		},
		{
			ValidatorName: "choices",
			Value:         1,
			Message:       "should be in [1, 2, 3]",
			Choices:       []int{1, 2, 3},
		},
	}

	// Valid cases
	for i, v := range vs {
		err := v.Validate(v.Value)
		if err != nil {
			t.Fatalf("case %d: expect: %v, got: %v", i, nil, err)
		}
	}

	// Invalid cases
	for i, v := range vs {
		err := v.Validate(i + 100)
		if err == nil {
			t.Fatalf("case %d: expect: %v, got: %v", i, v.Message, err)
		}
	}
}

func TestChoiceInvalidType(t *testing.T) {
	v := &ChoiceValidator{
		ValidatorName: "choices",
		Value:         "1",
		Message:       "should be in \"1\" or \"2\" or \"3\"",
		Choices:       "123",
	}

	defer func() {
		if e := recover().(string); e == "" {
			t.Fatalf(assertTemplate, errChoicesNotSlice, "")
		}
	}()

	v.Validate(v.Value)
}

func TestNewChoiceValidator(t *testing.T) {
	v := NewChoiceValidator("choices", "invalid", 1, []int{1, 2, 3})
	err := v.Validate(v.Value)
	if err != nil {
		t.Fatalf("expect: %v, got: %v", nil, err)
	}
}

func TestNewChoiceValidatorInvalidChoices(t *testing.T) {
	defer func() {
		if e := recover().(string); e == "" {
			t.Fatalf(assertTemplate, errChoicesNotSlice, "")
		}
	}()

	NewChoiceValidator("choices", "invalid", 1, "123")
}

func TestRemain(t *testing.T) {
	v := NewChoiceValidator("choices", "invalid", 1, []int{1, 2, 3})
	if v.Identifier() != "choices" {
		t.Fatalf(assertTemplate, "choices", v.Identifier())
	}

	v.SetValue(2)
	if !reflect.DeepEqual(v.Value, 2) {
		t.Fatalf(assertTemplate, 2, v.Value)
	}

	if !reflect.DeepEqual(v.GetValue(), 2) {
		t.Fatalf(assertTemplate, 2, v.GetValue())
	}
}
