package main

import (
	"sort"
)

func searchQuadruplets(a []int, t int) [][]int {
	sort.Ints(a)
	var res [][]int
	for i:=0; i < len(a)-3; i++ {
		if i > 0 && a[i] == a[i-1] {
			continue
		}
		for j:=i+1; j < len(a)-2; j++ {
			if j > i+1 && a[j] == a[j-1] {
				continue
			}
			searchPairs(a, t, i, j, &res)
		}
	}
	return res
}

func searchPairs(arr []int, t, f, s int, res *[][]int) {
	l := s+1
	r := len(arr)-1
	for l < r {
		sum := arr[f] + arr[s] + arr[l] + arr[r]
		if sum == t {
			*res = append(*res, []int{arr[f], arr[s], arr[l], arr[r]})
			l++
			r--
			for l < r && arr[l] == arr[l-1] {
				l++
			}
			for l < r && arr[r] == arr[r+1] {
				r--
			}
		} else if sum < t {
			l++
		} else {
			r--
		}
	}
}
