package main

func shortestWindowSort(arr []int) int {
	lp, rp := 0, len(arr)-1
	for lp < rp {
		lv := arr[lp]
		rv := arr[rp]
		if rv > lv && rv > arr[rp-1] {
			rp--
		} else if lv < rv && lv < arr[lp+1] {
			lp++
		} else {
			return rp-lp+1
		}
	}
	return rp-lp
}
