package validations

func SliceIsPresent(value []interface{}) bool {
	return len(value) > 0
}

// SliceIsUnique checks if all elements in a slice are unique
// by applying a function to each element and comparing the result.
//
// Example:
//
//	SliceIsUnique([]string{"a", "b", "c"}, func(s string) interface{} { return s })
//	SliceIsUnique([]int{1, 2, 3}, func(i int) interface{} { return i })
//	SliceIsUnique([]struct{ A, B int }{{1, 2}, {3, 4}, {5, 6}}, func(s struct{ A, B int }) interface{} { return s.A })
//
// Note: The function f must return a comparable type.
// See https://golang.org/ref/spec#Comparison_operators
func SliceIsUnique[T comparable](value []T, f func(T) interface{}) bool {
	seen := make(map[interface{}]bool)
	for _, v := range value {
		if seen[f(v)] {
			return false
		}
		seen[f(v)] = true
	}
	return true
}
