package main

func pairTargetSum(a []int, t int) []int {
	sp, ep := 0, len(a) - 1
	res := []int{}
	for sp < ep {
		lv := a[sp]
		ev := a[ep]
		s := lv + ev
		if s == t {
			res = append(res, sp, ep)
			return res
		}
		if s < t {
			sp++
		} else {
			ep--
		}
	}
	return res
}
