package dun

// If Conditional operator
// is expr = true result firstValue
// is expr = false result secondValue
func If[T any](expr bool, firstValue, secondValue T) T {
	if expr {
		return firstValue
	}
	return secondValue
}
