package validators

import "encoding/json"

// ValidateFieldName validation field name
type ValidateFieldName string

// ValidationErrors errors mapping for the validateion
type ValidationErrors map[ValidateIdentifier]error

// ValidationResult validation result
type ValidationResult map[ValidateFieldName]ValidationErrors

// MarshalJSON serialize validation result to json
func (r ValidationResult) MarshalJSON() ([]byte, error) {
	res := make(map[ValidateFieldName]map[ValidateIdentifier]string)
	for filedName, valErrs := range r {
		for id, err := range valErrs {
			if _, ok := res[filedName]; ok {
				res[filedName][id] = err.Error()
			} else {
				res[filedName] = map[ValidateIdentifier]string{
					id: err.Error(),
				}
			}
		}
	}
	return json.Marshal(res)
}

// IsValid check if the result contains error or not
func (r ValidationResult) IsValid() bool {
	if len(r) == 0 {
		return true
	}

	count := 0
	for _, valErrs := range r {
		for range valErrs {
			count++
		}
	}

	return count == 0
}

// FirstError random return an error in the validation result
func (r ValidationResult) FirstError() error {
	for _, ves := range r {
		for _, err := range ves {
			return err
		}
	}

	return nil
}

// Errors return all errors in a list
func (r ValidationResult) Errors() []error {
	errors := make([]error, 0)
	for _, ves := range r {
		for _, err := range ves {
			errors = append(errors, err)
		}
	}
	return errors
}
