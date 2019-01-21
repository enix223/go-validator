package validators

// Validator request param validator
type Validator map[ValidateFieldName][]ValidateHandler

// Validate perform `Validate` for each field in FieldValidators
func (v Validator) Validate() ValidationResult {
	res := make(map[ValidateFieldName]ValidationErrors)
	for fieldName, validateHandler := range v {
		for _, h := range validateHandler {
			err := h.Validate(h.GetValue())
			if err != nil {
				if _, ok := res[fieldName]; ok {
					res[fieldName][h.Identifier()] = err
				} else {
					res[fieldName] = ValidationErrors{
						h.Identifier(): err,
					}
				}
			}
		}
	}

	return res
}
