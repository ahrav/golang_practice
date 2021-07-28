package main

import (
	"math"
	"sort"

	"github.com/ahrav/learngo/algos/utils"
)

func tripletSumCloseTarget(a []int, t int) int {
	sort.Ints(a)
	res := int(math.Inf(-1))
	ip, lp, rp := 0, 1, len(a)-1
	for ip <= len(a) - 3 {
		for lp < rp {
			iv := a[ip]
			lv := a[lp]
			rv := a[rp]
			s := iv + lv + rv
			if t - s == 0 {
				return s
			}
			if int(utils.Abs(t-s)) < int(utils.Abs(t-res)) {
				if s > res {
					res = s
				}
			}
			if s < t {
				lp++
			} else {
				rp--
			}
		}
		ip++
		lp, rp = ip+1, len(a)-1
	}
	return res
}
