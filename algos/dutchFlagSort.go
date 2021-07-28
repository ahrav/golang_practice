package main

func dutchFlagSort(a []int) []int {
	x, l, h := 0, 0, len(a)-1
	for x <= h {
		if a[x] == 0 {
			a[l], a[x] = a[x], a[l]
			l++
			x++
		} else if a[x] == 2 {
			a[h], a[x] = a[x], a[h]
			h--
		} else {
			x++
		}
	}
	return a
}
