package validators

import (
	"reflect"
	"testing"
)

func TestFunctionValidator(t *testing.T) {
	f := func(value interface{}) bool {
		return true
	}

	v := &FunctionValidator{
		ValidatorName: "func",
		Message:       "invalid",
		Value:         1,
		Func:          f,
	}

	// valid case
	err := v.Validate(v.Value)
	if err != nil {
		t.Fatalf("expect: %v, got: %v", nil, err)
	}

	v.Func = func(value interface{}) bool {
		return false
	}
	err = v.Validate(v.Value)
	if err == nil {
		t.Fatalf("expect: %v, got: %v", v.Message, err)
	}
}

func TestFuncValidatorRemain(t *testing.T) {
	v := NewValidateFunc("func", "invalid", 1, func(value interface{}) bool {
		return true
	})
	if v.Identifier() != "func" {
		t.Fatalf(assertTemplate, "func", v.Identifier())
	}

	v.SetValue(2)
	if !reflect.DeepEqual(v.Value, 2) {
		t.Fatalf(assertTemplate, 2, v.Value)
	}

	if !reflect.DeepEqual(v.GetValue(), 2) {
		t.Fatalf(assertTemplate, 2, v.GetValue())
	}
}
