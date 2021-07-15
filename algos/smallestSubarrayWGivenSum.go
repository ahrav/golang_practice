package main

import (
	"math"

	"github.com/ahrav/learngo/algos/utils"
)


func smallestSubarrayWGivenSum(s int, a []int) int {
	if len(a) == 0 {
		return -1
	}
	r := math.MaxInt64
	wsu, wst := 0, 0
	for we := range a {
		wsu += a[we]
		for wsu >= s {
			r = utils.Min(r, we - wst + 1)
			wsu -= a[wst]
			wst++
		}
	}
	if r != math.MaxInt64 {
		return r
	}
	return 0
}
