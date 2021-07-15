package main

import "github.com/ahrav/learngo/algos/utils"

func maxSubArrayOfSizeK(k int, a []int) int {
	if len(a) == 0 {
		return -1
	}
	m := 0
	wsu, wst := 0, 0
	for we := range a {
		wsu += a[we]
		if we >= k - 1 {
			m = utils.Max(m, wsu)
			wsu -= a[wst]
			wst += 1
		}
	}
	return m
}
