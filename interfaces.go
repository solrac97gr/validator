package validator

type EvaluableStruct interface {
	Validate(args ...interface{}) error
}

type Validator interface {
	Struct(s EvaluableStruct) error
}
