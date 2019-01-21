package validators

// ValidateIdentifier validate handler identifier
type ValidateIdentifier string

// ValidateHandler - request parameter validation handler
type ValidateHandler interface {
	// Validate validate the input value, if valid, then return nil as error,
	// otherwise, return the error for the validation failure reason
	Validate(value interface{}) error

	// Identifier get ValidateHandler identifier
	Identifier() ValidateIdentifier

	// SetValue set value to be validated
	SetValue(value interface{})

	// GetValue get the value to be validated
	GetValue() interface{}
}
