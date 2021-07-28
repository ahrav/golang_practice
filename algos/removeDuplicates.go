package main

func removeDuplicates(a []int) int {
	lp, rp := 0, 1
	for rp < len(a) {
		if a[lp] == a[rp] {
			a = append(a[:lp], a[rp+1:]...)
		} else {
			lp++
			rp++
		}
	}
	return len(a)
}

func removeDuplicates2(a []int) int {
	nd, i := 1, 1
	for i < len(a) {
		if a[nd - 1] != a[i] {
			a[nd] = a[i]
			nd++
		}
		i++
	}
	return nd
}
