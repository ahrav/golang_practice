package main

func missingNum(arr []int) int {
	l := len(arr)
	for i, v := range arr {
		if v == i + 1 {
			return i
		}else if v == l {
			arr[i], arr[v-1] = arr[v-1], arr[i]
		}else {
			arr[i], arr[v] = arr[v], arr[i]
		}
	}
	return -1
}
