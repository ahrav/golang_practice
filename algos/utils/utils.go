// Package utils is a utility package.
package utils

// Max returns the larger of x or y.
func Max(x, y int) int {
	if x < y {
			return y
	}
	return x
}

// Min returns the smaller of x or y.
func Min(x, y int) int {
	if x > y {
			return y
	}
	return x
}

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
