package main

func findMissingNums(arr []int) []int {
	var res []int
	var i int
	for i < len(arr) {
		j := arr[i] - 1
		if arr[i] != arr[j] {
			arr[i], arr[j] = arr[j], arr[i]
		} else {
			i++
		}
	}

	for i := range arr {
		if i+1 != arr[i] {
			res = append(res, i+1)
		}
	}
	return res
}
