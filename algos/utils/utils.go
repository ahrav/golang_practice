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

// Insert adds value into specified index in a slice.
func Insert(a []int, idx int, val int) []int {
	if len(a) == idx {
		return append(a, val)
	}
	a = append(a[:idx+1], a[idx:]...)
	a[idx] = val
	return a
}

// Remove element from a slice while retaining the order.
func Remove(s [][]int, i int) []int {
	val := s[i]
	_ = append(s[:i], s[i+1:]...)
	return val
}
