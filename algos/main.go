package main

import (
	"fmt"
)


func main() {
	k := 5
	arr := []float64{1, 3, 2, 6, -1, 4, 1, 8, 2}
	fmt.Println(findAverageOfSubarrays(k, arr))
	k2 := 2
	arr2 := []int{2, 3, 4, 1, 5}
	fmt.Println(maxSubArrayOfSizeK(k2, arr2))
	k3 := 7
	arr3 := []int{2, 1, 5, 2, 3, 2}
	fmt.Println(smallestSubarrayWGivenSum(k3, arr3))
	k4 := 2
	str := "araaci"
	fmt.Println(longestSubstringWKDistinct(str, k4))
	arr4 := []string{"A", "B", "C", "B", "B", "C"}
	fmt.Println(fruitsIntoBasket(arr4))
	str2 := "abccde"
	fmt.Println(nonRepeatSubstring(str2))
	str3 := "aabccbb"
	k5 := 2
	fmt.Println(lengthLongestSubstring(str3, k5))
	arr5 := []int{0, 1, 1, 0, 0, 0, 1, 1, 0, 1, 1}
	k6 := 2
	fmt.Println(lengthLongestSubArray(arr5, k6))
}
