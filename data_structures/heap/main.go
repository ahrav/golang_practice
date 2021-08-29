package main

import "fmt"

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

func main() {
	m := &MaxHeap{}
	n := &MinHeap{}
	mh := &MedianHeap{min: n, max: m}
	// buildHeap := []int{10, 20, 30, 5, 7, 9, 11, 15, 17, 22}
	buildHeap2 := []int{3, 1, 5, 4}
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
	nums := []int{1, 2, -1, 3, 5, 7}
	fmt.Println(slidingWindow(nums, 3))

	// for i := 0; i < 5; i++ {
	// 	m.Extract()
	// 	n.Extract()
	// 	// fmt.Println(m)
	// 	// fmt.Println(n)
	// }
}
