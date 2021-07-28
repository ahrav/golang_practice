package main

import (
	"sort"
)

func searchTriplet(a []int) [][]int {
	sort.Ints(a)
	ip, lp, rp := 0, 1, len(a) - 1
	var res [][]int
	for ip <= len(a) - 3 {
		if a[ip] == 0 {
			break
		}
		t := -a[ip]
		for lp < rp {
			var x []int
			lv := a[lp]
			rv := a[rp]
			s := lv + rv
			if s == t {
				x = append(x, a[ip], lv, rv)
				res = append(res, x)
				for a[rp] == rv {
					rp--
				}
			} else if s < t {
				lp++
			} else {
				rp--
			}
		}
		ip++
		lp, rp = ip+1, len(a) - 1
	}
	return res
}
