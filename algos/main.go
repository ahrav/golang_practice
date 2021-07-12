package main

import "fmt"

// Max returns the larger of x or y.
func Max(x, y int) int {
	if x < y {
			return y
	}
	return x
}

func main() {
	k := 5
	arr := []float64{1, 3, 2, 6, -1, 4, 1, 8, 2}
	fmt.Println(findAverageOfSubarrays(k, arr))
	k2 := 2
	arr2 := []int{2, 3, 4, 1, 5}
	fmt.Println(maxSubArrayOfSizeK(k2, arr2))
}

func findAverageOfSubarrays(k int, a []float64) []float64 {
	r := []float64{}
	wsu, wst := 0.0, 0
	for we := range a {
		wsu += a[we]
		if we >= k - 1 {
			r = append(r, wsu / float64(k))
			wsu -= a[wst]
			wst += 1
		}
	}

	return r
}

func maxSubArrayOfSizeK(k int, a []int) int {
	if len(a) == 0 {
		return -1
	}
	m := 0
	wsu, wst := 0, 0
	for we := range a {
		wsu += a[we]
		if we >= k - 1 {
			m = Max(m, wsu)
			wsu -= a[wst]
			wst += 1
		}
	}
	return m
}
