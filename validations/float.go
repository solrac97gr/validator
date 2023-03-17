package validations

func FloatIsPositive(value float64) bool {
	return value > 0
}

func FloatIsNegative(value float64) bool {
	return value < 0
}

func FloatIsZero(value float64) bool {
	return value == 0
}

func FloatIsNonZero(value float64) bool {
	return value != 0
}

func FloatIsInRange(value float64, min float64, max float64) bool {
	return value >= min && value <= max
}

func FloatIsLessThan(value float64, max float64) bool {
	return value < max
}

func FloatIsLessThanOrEqualTo(value float64, max float64) bool {
	return value <= max
}

func FloatIsGreaterThan(value float64, min float64) bool {
	return value > min
}

func FloatIsGreaterThanOrEqualTo(value float64, min float64) bool {
	return value >= min
}

func FloatIsMultipleOf(value float64, multiple float64) bool {
	return value/multiple == 0
}
