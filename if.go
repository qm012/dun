package dun

// If Conditional operator
// is expr = true result left
// is expr = false result right
func If[T any](expr bool, left, right T) T {
	if expr {
		return left
	}
	return right
}
