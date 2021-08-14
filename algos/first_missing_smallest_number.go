package main

func firstSmallestMissingNum(arr []int) int {
	var i int
	for i < len(arr) {
		j := arr[i] - 1
		if arr[i] > 0 && arr[i] < len(arr) && arr[i] != arr[j] {
			arr[i], arr[j] = arr[j], arr[i]
		} else {
			i++
		}
	}
	for k, v := range(arr) {
		if v != k+1 {
			return k+1
		}
	}
	return len(arr) + 1
}
