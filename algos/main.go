package main

import "fmt"


func main() {
	// k := 5
	// arr := []float64{1, 3, 2, 6, -1, 4, 1, 8, 2}
	// fmt.Println(findAverageOfSubarrays(k, arr))
	// k2 := 2
	// arr2 := []int{2, 3, 4, 1, 5}
	// fmt.Println(maxSubArrayOfSizeK(k2, arr2))
	// k3 := 7
	// arr3 := []int{2, 1, 5, 2, 3, 2}
	// fmt.Println(smallestSubarrayWGivenSum(k3, arr3))
	// k4 := 2
	// str := "araaci"
	// fmt.Println(longestSubstringWKDistinct(str, k4))
	// arr4 := []string{"A", "B", "C", "B", "B", "C"}
	// fmt.Println(fruitsIntoBasket(arr4))
	// str2 := "abccde"
	// fmt.Println(nonRepeatSubstring(str2))
	// str3 := "aabccbb"
	// k5 := 2
	// fmt.Println(lengthLongestSubstring(str3, k5))
	// arr5 := []int{0, 1, 1, 0, 0, 0, 1, 1, 0, 1, 1}
	// k6 := 2
	// fmt.Println(lengthLongestSubArray(arr5, k6))
	// str4 := "oidbcaf"
	// p := "abc"
	// fmt.Println(stringPermutations2(str4, p))
	// str5 := "ppqp"
	// p2 := "pq"
	// fmt.Println(stringAnagram(str5, p2))
	// str6 := "aabdec"
	// p3 := "abc"
	// fmt.Println(shortestSubstring(str6, p3))
	// str7 := "catfoxcat"
	// w := []string{"cat", "fox"}
	// fmt.Println(wordConcat(str7, w))
	// i := []int{1, 2, 3, 4, 6}
	// t := 6
	// fmt.Println(pairTargetSum(i, t))
	// i2 := []int{2, 3, 3, 3, 6, 9, 9}
	// fmt.Println(removeDuplicates(i2))
	// i3 := []int{-3, -1, 0, 1, 2}
	// fmt.Println(makeSquares(i3))
	// i4 := []int{-3, 0, 1, 2, -1, 1, -2}
	// fmt.Println(searchTriplet(i4))
	// arr6 := []int{1, 0, 1, 1}
	// t := 100
	// fmt.Println(tripletSumCloseTarget(arr6, t))
	// arr7 := []int{-1, 0, 2, 3}
	// t2 := 3
	// fmt.Println(tripletSmallerSum(arr7, t2))
	// arr8 := []int{2, 5, 3, 10}
	// t3 := 30
	// fmt.Println(findSubarray(arr8, t3))
	// arr9 := []int{1, 0, 2, 1, 0}
	// fmt.Println(dutchFlagSort(arr9))
	// arr10 := []int{2, 0, -1, 1, -2, 2}
	// t4 := 2
	// fmt.Println(searchQuadruplets(arr10, t4))
	// fmt.Println(backspaceCompare("xp#", "xyz##"))
	// arr11 := []int{1, 3, 2, 0, -1, 7, 10}
	// fmt.Println(shortestWindowSort(arr11))
	// mergeIntervalsExample()
	// cyclicArr := []int{2, 6, 4, 3, 1, 5}
	// fmt.Println(cyclicSort(cyclicArr))
	// missingNumArr := []int{8, 3, 5, 2, 4, 6, 0, 1}
	// fmt.Println(missingNum(missingNumArr))
	// missingNumsArr := []int{2, 3, 1, 8, 2, 3, 5, 1}
	// missingNumsArr2 := []int{2, 4, 1, 2}
	// missingNumsArr3 := []int{2, 3, 2, 1}
	// fmt.Println(findMissingNums(missingNumsArr))
	// fmt.Println(findMissingNums(missingNumsArr2))
	// fmt.Println(findMissingNums(missingNumsArr3))
	// duplicateArr := []int{1, 4, 4, 3, 2}
	// duplicateArr2 := []int{2, 1, 3, 3, 5, 4}
	// duplicateArr3 := []int{2, 4, 1, 4, 4}
	// fmt.Println(findDuplicates(duplicateArr))
	// fmt.Println(findDuplicates(duplicateArr2))
	// fmt.Println(findDuplicates(duplicateArr3))
	// allDuplictesArr := []int{3, 4, 4, 5, 5}
	// allDuplictesArr2 := []int{5, 4, 7, 2, 3, 5, 3}
	// fmt.Println(findAllDuplicates(allDuplictesArr))
	// fmt.Println(findAllDuplicates(allDuplictesArr2))
	// corruptNumArr := []int{3, 1, 2, 5, 2}
	// corruptNumArr2 := []int{3, 1, 2, 3, 6, 4}
	// fmt.Println(findCorruptNumbers(corruptNumArr))
	// fmt.Println(findCorruptNumbers(corruptNumArr2))
	// firstSmallestMissingArr := []int{3, -2, 0, 1, 2}
	// firstSmallestMissingArr2 := []int{3, 2, 5, 1}
	// fmt.Println(firstSmallestMissingNum(firstSmallestMissingArr))
	// fmt.Println(firstSmallestMissingNum(firstSmallestMissingArr2))
	kMissingPositiveArr := []int{3, -1, 4, 5, 5}
	kMissingPositiveArr2 := []int{2, 3, 4}
	kMissingPositiveArr3 := []int{-2, -3, 4}
	fmt.Println(firstKMissingPositive2(kMissingPositiveArr, 3))
	fmt.Println(firstKMissingPositive2(kMissingPositiveArr2, 3))
	fmt.Println(firstKMissingPositive2(kMissingPositiveArr3, 2))
}
