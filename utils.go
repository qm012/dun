package ok

// If Conditional operator
// is expr = true return firstValue
// is expr = false return secondValue
func If[T any](expr bool, firstValue, secondValue T) T {
	if expr {
		return firstValue
	}
	return secondValue
}
