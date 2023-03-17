package validations

func IsNegative[T int](value T) bool {
	return value < 0
}

func IsZero[T int](value T) bool {
	return value == 0
}

func IsNonZero[T int](value T) bool {
	return value != 0
}

func IsEven[T int](value T) bool {
	return value%2 == 0
}

func IsOdd[T int](value T) bool {
	return value%2 != 0
}

func IsDivisibleBy[T int](value T, divisor T) bool {
	return value%divisor == 0
}

func IsInRange[T int](value T, min T, max T) bool {
	return value >= min && value <= max
}

func IsLessThan[T int](value T, max T) bool {
	return value < max
}

func IsLessThanOrEqualTo[T int](value T, max T) bool {
	return value <= max
}

func IsGreaterThan[T int](value T, min T) bool {
	return value > min
}

func IsGreaterThanOrEqualTo[T int](value T, min T) bool {
	return value >= min
}

func IsMultipleOf[T int](value T, multiple T) bool {
	return value%multiple == 0
}
