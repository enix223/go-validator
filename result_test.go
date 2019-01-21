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
