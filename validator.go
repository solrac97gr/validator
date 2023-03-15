package validator

// ValidatorImpl is the default implementation of the Validator interface.
type ValidatorImpl struct {
}

// The Validator interface is implemented by ValidatorImpl.
var _ Validator = &ValidatorImpl{}

// NewValidator returns a new ValidatorImpl.
func NewValidator() *ValidatorImpl {
	return &ValidatorImpl{}
}

// Struct validates the given struct.
func (v *ValidatorImpl) Struct(s EvaluableStruct) error {
	return s.Validate()
}
