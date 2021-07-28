package main

import (
	"math"

	"github.com/ahrav/learngo/algos/utils"
)

func makeSquares(a []int) []int {
	var res []int
	lp, rp := 0, len(a) - 1
	for lp <= rp {
		if utils.Abs(a[lp]) > utils.Abs(a[rp]) {
			v := math.Pow(float64(a[lp]), 2)
			res = append([]int{int(v)}, res...)
			lp++
		} else {
			v := math.Pow(float64(a[rp]), 2)
			res = append([]int{int(v)}, res...)
			rp--
		}
	}
	return res
}
