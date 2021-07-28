package main

import (
	"sort"
)

func tripletSmallerSum(a []int, t int) int {
	res := 0
	sort.Ints(a)
	ip, lp, rp := 0, 1, len(a)-1
	for ip <= len(a)-3 {
		for lp < rp {
			iv := a[ip]
			lv := a[lp]
			rv := a[rp]
			s := rv + lv + iv
			if s < t {
				res+=rp-lp
				lp++
			} else {
				rp--
			}
		}
		ip++
		lp= ip+1
	}
	return res
}
