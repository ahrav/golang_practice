package main

import "fmt"

// MaxHeap struct has slice that holds array.
type MaxHeap struct{
	array []int
}

// Insert adds an element to a heap.
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array)-1)
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

// maxHeapifyUp will heapify from the bottom up.
func (h *MaxHeap) maxHeapifyUp(idx int) {
	for h.array[parent(idx)] < h.array[idx] {
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

func main() {
	m := &MaxHeap{}
	buildHeap := []int{10, 20, 30, 5, 7, 9, 11, 15, 17, 22}
	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}

	for i := 0; i < 5; i++ {
		m.Extract()
		fmt.Println(m)
	}
}
