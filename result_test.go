package validators

import (
	"encoding/json"
	"errors"
	"testing"
)

func TestValidatorResult(t *testing.T) {
	res := make(ValidationResult)

	res["abc"] = make(ValidationErrors)
	res["abc"]["e1"] = errors.New("err1")
	res["abc"]["e2"] = errors.New("err2")

	bytes, err := json.Marshal(res)
	if err != nil {
		t.Fatalf(assertTemplate, nil, err)
	}

	expect := `{"abc":{"e1":"err1","e2":"err2"}}`
	if string(bytes) != expect {
		t.Fatalf(assertTemplate, expect, string(bytes))
	}
}

func TestBlankValidatorResult(t *testing.T) {
	var res ValidationResult
	if !res.IsValid() {
		t.Fatalf(assertTemplate, true, false)
	}
}

func TestNonBlankValidatorResult(t *testing.T) {
	res := make(ValidationResult)

	res["abc"] = make(ValidationErrors)
	res["abc"]["e"] = errors.New("err")

	if res.IsValid() {
		t.Fatalf(assertTemplate, false, true)
	}
}

func TestErrors(t *testing.T) {
	res := make(ValidationResult)

	res["abc"] = make(ValidationErrors)
	res["abc"]["e1"] = errors.New("err1")
	res["abc"]["e2"] = errors.New("err2")
	res["abb"] = make(ValidationErrors)
	res["abb"]["e3"] = errors.New("err3")
	res["abb"]["e4"] = errors.New("err4")

	errors := []error{
		res["abc"]["e1"],
		res["abc"]["e2"],
		res["abb"]["e3"],
		res["abb"]["e4"],
	}

	for _, err := range res.Errors() {
		if !containError(err, errors) {
			t.Fatalf("Errors not correct")
		}
	}
}

func containError(err error, errors []error) bool {
	var found bool
	for _, er := range errors {
		if er == err {
			found = true
			break
		}
	}
	return found
}

func TestFirstError(t *testing.T) {
	res := make(ValidationResult)

	res["abc"] = make(ValidationErrors)
	res["abc"]["e1"] = errors.New("err1")
	res["abc"]["e2"] = errors.New("err2")
	res["abb"] = make(ValidationErrors)
	res["abb"]["e3"] = errors.New("err3")
	res["abb"]["e4"] = errors.New("err4")

	errors := []error{
		res["abc"]["e1"],
		res["abc"]["e2"],
		res["abb"]["e3"],
		res["abb"]["e4"],
	}

	if !containError(res.FirstError(), errors) {
		t.Fatalf("FirstError not correct")
	}

	// test empty error
	res = make(ValidationResult)
	if res.FirstError() != nil {
		t.Fatalf("FirstError not correct")
	}
}
