package main

func findCorruptNumbers(arr []int) []int {
	var res []int
	var i int
	for i < len(arr) {
		j := arr[i]-1
		if arr[i] != arr[j] {
			arr[i], arr[j] = arr[j], arr[i]
		} else {
			i++
		}
	}
	for k, v := range arr {
		if v != k+1 {
			res = append(res, k+1, v)
		}
	}
	return res
}
