package main

func cyclicSort(arr []int) []int {
	for i, val := range arr {
		if val == i+1 {
			continue
		}
		arr[i], arr[val-1] = arr[val-1], arr[i]
	}
	return arr
}
