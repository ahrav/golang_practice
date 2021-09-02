package main

import (
	"container/heap"
	"fmt"

	"github.com/ahrav/learngo/algos/utils"
)

// MinH is a min-heap of ints.
type MinH [][]int

func (h MinH) Len() int           { return len(h) }
func (h MinH) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h MinH) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinH) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.([]int))
}

func (h *MinH) Pop() []int {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *MinH) Peek() []int {
	old := *h
	return old[0]
}

// MaxH is a max-heap of ints.
type MaxH [][]int

func (h MaxH) Len() int           { return len(h) }
func (h MaxH) Less(i, j int) bool { return h[i][0] > h[j][0] }
func (h MaxH) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxH) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.([]int))
}

func (h *MaxH) Pop() []int {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *MaxH) Peek() []int {
	old := *h
	return old[0]
}

// MaxStartHeap is a max-heap of ints.
type MaxStartHeap [][]int

func (h MaxStartHeap) Len() int           { return len(h) }
func (h MaxStartHeap) Less(i, j int) bool { return h[i][0] > h[j][0] }
func (h MaxStartHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxStartHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.([]int))
}

func (h *MaxStartHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *MaxStartHeap) Peek() []int {
	old := *h
	return old[0]
}

// MaxEndHeap is a max-heap of ints.
type MaxEndHeap [][]int

func (h MaxEndHeap) Len() int           { return len(h) }
func (h MaxEndHeap) Less(i, j int) bool { return h[i][0] > h[j][0] }
func (h MaxEndHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxEndHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.([]int))
}

func (h *MaxEndHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *MaxEndHeap) Peek() []int {
	old := *h
	return old[0]
}

// MaxHeap struct has slice that holds array.
type MaxHeap struct{
	array []int
}

// MinHeap struct has slice that holds array.
type MinHeap struct{
	array []int
}

// MedianHeap struct has a min and max heap.
type MedianHeap struct{
	min *MinHeap
	max *MaxHeap
}

// Insert adds an element to a heap.
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array)-1)
}

// Insert adds an element to a heap.
func (h *MinHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array)-1)
}

// Insert adds an element to a heap.
func (h *MedianHeap) Insert(key int) {
	if len(h.max.array) == 0 && len(h.min.array) == 0 {
		h.min.array = append(h.min.array, key)
		h.min.maxHeapifyUp(len(h.min.array)-1)
		return
	}
	median := int(h.Median())
	if key >= median {
		h.min.array = append(h.min.array, key)
		h.min.maxHeapifyUp(len(h.min.array)-1)
	} else {
		h.max.array = append(h.max.array, key)
		h.max.maxHeapifyUp(len(h.max.array)-1)
	}
	diff := len(h.min.array) - len(h.max.array)
	if diff > 1 {
		h.max.Insert(h.min.Extract())
	} else if diff < -1 {
		h.min.Insert(h.max.Extract())
	}
}

// Median returns median value of the heap.
func (h *MedianHeap) Median() float64 {
	if len(h.min.array) > len(h.max.array) {
		return float64(h.min.array[0])
	} else if len(h.min.array) < len(h.max.array) {
		return float64(h.max.array[0])
	} else {
		return float64((h.max.array[0] + h.min.array[0])) / 2.0
	}
}

// Extract returns the largest key, and removes it from the heap.
func (h *MaxHeap) Extract() int {
	ext := h.array[0]
	l := len(h.array)-1

	if len(h.array) == 0 {
		fmt.Println("Can not extract because array is length 0")
		return -1
	}
	h.array[0] = h.array[l]
	h.array = h.array[:l]

	h.maxHeapifyDown(0)
	return ext
}

// Extract returns the smallest key, and removes it from the heap.
func (h *MinHeap) Extract() int {
	ext := h.array[0]
	l := len(h.array)-1

	if len(h.array) == 0 {
		fmt.Println("Can not extract because array is length 0")
		return -1
	}
	h.array[0] = h.array[l]
	h.array = h.array[:l]

	h.maxHeapifyDown(0)
	return ext
}

// maxHeapifyUp will heapify from the bottom up.
func (h *MaxHeap) maxHeapifyUp(idx int) {
	for h.array[parent(idx)] < h.array[idx] {
		h.swap(parent(idx), idx)
		idx = parent(idx)
	}
}

// maxHeapifyUp will heapify from the bottom up.
func (h *MinHeap) maxHeapifyUp(idx int) {
	for h.array[parent(idx)] > h.array[idx] {
		h.swap(parent(idx), idx)
		idx = parent(idx)
	}
}

// maxHeapifyDown will heapify from the top to bottom.
func (h *MaxHeap) maxHeapifyDown(idx int) {
	lastIdx := len(h.array)-1
	l,r := left(idx), right(idx)
	childCompare := 0

	for l <= lastIdx {
		if l == lastIdx { // left child is the only child.
			childCompare = l
		} else if h.array[l] > h.array[r] { // left child is larger
			childCompare = l
		} else { // right child is larger
			childCompare = r
		}
		// compare array value of current index to larger child and swap i smaller
		if h.array[idx] < h.array[childCompare] {
			h.swap(idx, childCompare)
			idx = childCompare
			l, r = left(idx), right(idx)
		} else {
			return
		}
	}
}

// maxHeapifyDown will heapify from the top to bottom.
func (h *MinHeap) maxHeapifyDown(idx int) {
	lastIdx := len(h.array)-1
	l,r := left(idx), right(idx)
	childCompare := 0

	for l <= lastIdx {
		if l == lastIdx { // left child is the only child.
			childCompare = l
		} else if h.array[l] < h.array[r] { // left child is smaller
			childCompare = l
		} else { // right child is smaller
			childCompare = r
		}
		// compare array value of current index to smaller child and swap i larger
		if h.array[idx] > h.array[childCompare] {
			h.swap(idx, childCompare)
			idx = childCompare
			l, r = left(idx), right(idx)
		} else {
			return
		}
	}
}

// get the parent index.
func parent(i int) int {
	return (i-1)/2
}

// get left child index.
func left(i int) int {
	return 2*i + 1
}

// get right child index.
func right(i int) int {
	return 2*i + 2
}

// swap keys in the array.
func (h *MaxHeap) swap (i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

// swap keys in the array.
func (h *MinHeap) swap (i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func slidingWindow(arr []int, k int) []float64 {
	var res []float64
	h := MedianHeap{min: &MinHeap{}, max: &MaxHeap{}}
	for idx, v := range arr {
		h.Insert(v)
		if idx >= k-1 {
			res = append(res, h.Median())
			h = MedianHeap{min: &MinHeap{}, max: &MaxHeap{}}
			for i:=0; i<k-1; i++ {
				h.Insert(arr[idx-i])
			}
		}
	}
	return res
}

func maxCapital(c, p []int, ic, n int) int {
	res := ic
	for n > 0 {
		var max int
		for i, e := range c {
			if res >= e {
				max = utils.Max(max, p[i])
			}
		}
		res+=max
		n--
		max = 0
	}
	return res
}

func maxCapitalHeap(c, p []int, ic, n int) int {
	min := MinH{}
	max := MaxH{}
	for i, e := range c {
		min.Push([]int{e, i})
	}
	cap := ic
	for i:=0; i<n; i++ {
		for min.Len() > 0 && min.Peek()[0] <= cap {
			max.Push(p[min.Pop()[1]])
		}
		if max.Len() == 0 {
			break
		}
		cap+=max.Pop()[0]
	}
	return cap
}

func nextInterval(interval [][]int) []int {
	min := &MaxStartHeap{}
	max := &MaxEndHeap{}
	heap.Init(min)
	heap.Init(max)
	res := make([]int, len(interval))
	for i, e := range interval {
		heap.Push(min, []int{e[0], i})
		heap.Push(max, []int{e[1], i})
	}
	for range interval {
		topEnd := heap.Pop(max)
		end := topEnd.([]int)
		res[end[1]] = -1
		if min.Peek()[0] >= end[0] {
			start := heap.Pop(min)
			for min.Len() > 0 && min.Peek()[0] >= end[0] {
				start = heap.Pop(min)
			}
			tStart := start.([]int)
			res[end[1]] = tStart[1]
			heap.Push(min, start)
		}
	}
	return res
}

func main() {
	// fmt.Println(maxCapital([]int{0, 1, 2, 3}, []int{1, 2, 3, 5}, 0, 3))
	m := &MaxHeap{}
	n := &MinHeap{}
	mh := &MedianHeap{min: n, max: m}
	// buildHeap := []int{10, 20, 30, 5, 7, 9, 11, 15, 17, 22}
	buildHeap2 := []int{3, 1, 5, 4}
	fmt.Println(nextInterval([][]int{{2, 3}, {3, 4}, {5, 6}}))
	// for _, v := range buildHeap {
	// 	// m.Insert(v)
	// 	// n.Insert(v)
	// 	mh.Insert(v)
	// 	// fmt.Println(m)
	// 	// fmt.Println(n)
	// 	// fmt.Println(mh)
	// }
	for _, v := range buildHeap2 {
		mh.Insert(v)
		// fmt.Println(m)
		// fmt.Println(n)
		// fmt.Println(mh)
	}
	// fmt.Println(mh.Median())
	// nums := []int{1, 2, -1, 3, 5, 7}
	// fmt.Println(slidingWindow(nums, 3))

	// for i := 0; i < 5; i++ {
	// 	m.Extract()
	// 	n.Extract()
	// 	// fmt.Println(m)
	// 	// fmt.Println(n)
	// }
}
