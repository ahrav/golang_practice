package main

func findAllDuplicates(arr []int) []int {
	var res []int
	var i int
	for i < len(arr) {
		j := arr[i]-1
		if arr[i] != i+1 {
			if arr[i] != arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			} else {
				res = append(res, arr[i])
				i++
			}
		} else {
			i++
		}
	}
	return res
}
