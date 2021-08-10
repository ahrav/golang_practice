package main

import (
	"container/heap"
)

// An IntHeap is a min-heap of ints.
type intHeap []*Interval

type employeeInterval struct {
	Interval
	idx int
	intervalIdx int
}

func (h intHeap) Len() int           { return len(h) }
func (h intHeap) Less(i, j int) bool { return h[i].Start < h[j].Start }
func (h intHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *intHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(*Interval))
}

func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func freeTimeInterval(schedule [][]*Interval) []*Interval {
	var intHeap intHeap
	heap.Init(&intHeap)

	// Flatten array
	var flat []*Interval
	for i:=0; i<len(schedule); i++ {
		flat = append(flat, schedule[i]...)
	}

	for i:=0; i<len(flat); i++ {
		heap.Push(&intHeap, flat[i])
	}

	temp := heap.Pop(&intHeap).(*Interval)
	var res []*Interval

	for intHeap.Len() > 0 {
		a := temp
		b := heap.Pop(&intHeap).(*Interval)

		if intervalOverlap(a, b) {
			if a.End < b.End {
				temp = b
			}
		} else {
			res = append(res, &Interval{Start: a.End, End: b.Start})
			temp = b
		}
	}
	return res
}

func intervalOverlap(a, b *Interval) bool {
	if a.Start > b.End {
		return false
	}
	if b.Start > a.End {
		return false
	}
	return true
}
