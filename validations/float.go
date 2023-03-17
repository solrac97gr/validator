package validations

func FloatIsPositive[T float64](value T) bool {
	return value > 0
}

func FloatIsNegative[T float64](value T) bool {
	return value < 0
}

func FloatIsZero[T float64](value T) bool {
	return value == 0
}

func FloatIsNonZero[T float64](value T) bool {
	return value != 0
}

func FloatIsInRange[T float64](value T, min T, max T) bool {
	return value >= min && value <= max
}

func FloatIsLessThan[T float64](value T, max T) bool {
	return value < max
}

func FloatIsLessThanOrEqualTo[T float64](value T, max T) bool {
	return value <= max
}

func FloatIsGreaterThan[T float64](value T, min T) bool {
	return value > min
}

func FloatIsGreaterThanOrEqualTo[T float64](value T, min T) bool {
	return value >= min
}

func FloatIsMultipleOf[T float64](value T, multiple T) bool {
	return value/multiple == 0
}
