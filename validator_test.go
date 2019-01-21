package validators

import "testing"

func TestValidValidator(t *testing.T) {
	val := Validator{
		"abc": []ValidateHandler{
			&BlankValidator{
				Value:         "1",
				Message:       "err",
				BlankValue:    "0",
				ValidatorName: "blank",
			},
		},
	}

	res := val.Validate()
	if !res.IsValid() {
		t.Fatalf(assertTemplate, true, false)
	}
}

func TestValidMultiValidator(t *testing.T) {
	val := Validator{
		"abc": []ValidateHandler{
			&BlankValidator{
				Value:         "1",
				Message:       "err",
				BlankValue:    "0",
				ValidatorName: "blank",
			},
			&BlankValidator{
				Value:         "2",
				Message:       "err",
				BlankValue:    "0",
				ValidatorName: "blank",
			},
		},
		"bbc": []ValidateHandler{
			&BlankValidator{
				Value:         "1",
				Message:       "err",
				BlankValue:    "0",
				ValidatorName: "blank",
			},
			&BlankValidator{
				Value:         "2",
				Message:       "err",
				BlankValue:    "0",
				ValidatorName: "blank",
			},
		},
	}

	res := val.Validate()
	if !res.IsValid() {
		t.Fatalf(assertTemplate, true, false)
	}
}

func TestInValidValidator(t *testing.T) {
	val := Validator{
		"abc": []ValidateHandler{
			&BlankValidator{
				Value:         "0",
				Message:       "err",
				BlankValue:    "0",
				ValidatorName: "blank",
			},
		},
	}

	res := val.Validate()
	if res.IsValid() {
		t.Fatalf(assertTemplate, false, true)
	}
}
