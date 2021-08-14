package main

func findDuplicates(arr []int) int {
	var i int
	for i < len(arr) {
		j := arr[i] - 1
		if arr[i] != i + 1 {
			if arr[i] != arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			} else {
				return arr[i]
			}
		} else {
			i++
		}
	}
	return -1
}
