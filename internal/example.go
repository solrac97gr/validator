package example

import (
	"errors"

	"github.com/solrac97gr/validator"
)

// Square is a struct that represents a square.
type Square struct {
	Side1Length float64
	Side2Length float64
	Side3Length float64
	Side4Length float64
}

// The EvaluableStruct interface is implemented by Square.
var _ validator.EvaluableStruct = &Square{}

// Validate validates the square.
func (s Square) Validate(args ...interface{}) error {
	if s.Side1Length != s.Side2Length || s.Side2Length != s.Side3Length || s.Side3Length != s.Side4Length {
		return errors.New("not a square - sides are not equal")
	}
	return nil
}

// CreateSquare creates a new square.
func CreateSquare(sideLength float64, val validator.Validator) *Square {
	sq := new(Square)
	sq.Side1Length = sideLength
	sq.Side2Length = sideLength
	sq.Side3Length = sideLength
	sq.Side4Length = sideLength

	if err := val.Struct(sq); err != nil {
		return nil
	}

	return sq
}
