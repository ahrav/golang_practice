package main

import "sort"

func firstKMissingPositive(arr []int, k int) []int {
	sort.Ints(arr)
	e := 1
	var i int
	var res []int
	for k > 0 {
		if i < len(arr) && arr[i] > 0 {
			if arr[i] == e {
				e++
				i++
			} else {
				res = append(res, e)
				e++
				k--
			}
		} else {
			i++
			res = append(res, e)
			e++
			k--
		}
	}
	return res
}

func firstKMissingPositive2(arr []int, k int) []int {
	var i int
	for i < len(arr) {
		j := arr[i] - 1
		if arr[i] > 0 && arr[i] <= len(arr) && arr[i] != arr[j] {
			arr[i], arr[j] = arr[j], arr[i]
		} else {
			i++
		}
	}

	var missingNums []int
	m := map[int]bool{}
	for key, v := range(arr) {
		if len(missingNums) < k {
			if v != key+1 {
				missingNums = append(missingNums, key+1)
				m[v] = true
			}
		}
	}
	i = 1
	for len(missingNums) < k {
		num := i + len(arr)
		if _, ok := m[num]; ok {
			i++
		} else {
			missingNums = append(missingNums, num)
		}
	}
	return missingNums
}
