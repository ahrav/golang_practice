package main

import (
	"fmt"
	"sort"

	"github.com/ahrav/learngo/algos/utils"
)

// Interval struct
type Interval struct {
	Start int
	End int
}

type job struct {
	Load int
	Interval
}

func mergeIntervals(intervals []Interval) []Interval {
	if len(intervals) < 2 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	var mergedIntervals []Interval
	s := intervals[0].Start
	e := intervals[0].End
	for i:=1; i < len(intervals); i++ {
		interval := intervals[i]
		if interval.Start <= e {
			e = utils.Max(interval.End, e)
		} else {
			mergedIntervals = append(mergedIntervals, Interval{Start: s, End: e})
			s = interval.Start
			e = interval.End
		}
	}

	mergedIntervals = append(mergedIntervals, Interval{Start: s, End: e})

	return mergedIntervals
}

func doIntervalsOverlap(intervals []Interval) bool {
	if len(intervals) < 2 {
		return false
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	e := intervals[0].End
	for i:=1; i < len(intervals); i++ {
		interval := intervals[i]
		if interval.Start <= e {
			return true
		} else {
			e = interval.End
		}
	}

	return false
}

func insertMerged(intervals []Interval, newInterval Interval) []Interval {
	if len(intervals) < 1 {
		return intervals
	}

	var mergedIntervals []Interval
	i := 0

	for i < len(intervals) && intervals[i].End < newInterval.Start {
		mergedIntervals = append(mergedIntervals, intervals[i])
		i++
	}

	for i < len(intervals) && intervals[i].Start <= newInterval.End {
		newInterval.Start = utils.Min(intervals[i].Start, newInterval.Start)
		newInterval.End = utils.Max(intervals[i].End, newInterval.End)
		i++
	}
	mergedIntervals = append(mergedIntervals, newInterval)

	for i < len(intervals) {
		mergedIntervals = append(mergedIntervals, intervals[i])
		i++
	}

	return mergedIntervals
}

func mergeIntersection(intervals1 []Interval, intervals2 []Interval) []Interval {
	if len(intervals1) == 0 {
		return nil
	}
	if len(intervals2) == 0 {
		return nil
	}
	var unionIntervals []Interval

	i, j, il1, il2 := 0, 0, len(intervals1), len(intervals2)
	for i < il1 && j < il2 {
		i1 := intervals1[i]
		i2 := intervals2[j]
		i1OverI2 := i1.Start >= i2.Start && i1.Start <= i2.End
		i2OverI1 := i2.Start >= i1.Start && i2.Start <= i1.End
		if i1OverI2 || i2OverI1 {
			unionIntervals = append(unionIntervals, Interval{Start: utils.Max(i1.Start, i2.Start), End: utils.Min(i1.End, i2.End)})
		}

		if i1.End > i2.End {
			j++
		} else {
			i++
		}
	}
	return unionIntervals
}

func hasOverlap(intervals []Interval) bool {
	if len(intervals) <= 1 {
		return false
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	e := intervals[0].End
	for i:=1; i<len(intervals); i++ {
		interval := intervals[i]
		if interval.Start < e {
			return true
		}
		e = interval.End
	}
	return false
}

func overlapIntervals(intervals []Interval) []Interval {
	if len(intervals) <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	var overlaps []Interval
	s := intervals[0].Start
	e := intervals[0].End
	for i:=1; i<len(intervals); i++ {
		interval := intervals[i]
		if interval.Start < e {
			overlaps = append(overlaps, Interval{Start: s, End: e}, Interval{Start: interval.Start, End: interval.End})
			continue
		}
		e = interval.End
		s = interval.Start
	}
	return overlaps
}

func minMeetings(intervals []Interval) int {
	if len(intervals) <= 1 {
		return 1
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	c := 1
	e := intervals[0].End
	for i:=1; i<len(intervals); i++ {
		interval := intervals[i]
		if interval.Start < e {
			c++
			e = interval.End
			continue
		}
		e = interval.End
	}
	return c
}

func maxLoad(jobs []job) int {
	if len(jobs) == 0 {
		return -1
	}
	if len(jobs) == 1 {
		return jobs[0].Load
	}

	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i].Start < jobs[j].Start
	})

	l := jobs[0].Load
	e := jobs[0].End
	for i:=1; i<len(jobs); i++ {
		job := jobs[i]
		if job.Start < e {
			l+=job.Load
			continue
		}
		l = utils.Max(job.Load, l)
	}
	return l
}



func mergeIntervalsExample() {
	intervals := []Interval{{Start: 6, End: 7}, {Start: 2, End: 4}, {Start: 5, End: 9}}
	fmt.Println(mergeIntervals(intervals))
	intervals2 := []Interval{{Start: 6, End: 7}, {Start: 8, End: 9}, {Start: 11, End: 19}}
	fmt.Println(doIntervalsOverlap(intervals2))
	intervals3 := []Interval{{Start: 1, End: 3}, {Start: 5, End: 7}, {Start: 8, End: 12}}
	fmt.Println(insertMerged(intervals3, Interval{Start: 4, End: 6}))
	intervals4 := []Interval{{Start: 1, End: 3}, {Start: 5, End: 7}, {Start: 9, End: 12}}
	intervals5 := []Interval{{Start: 5, End: 10}}
	fmt.Println(mergeIntersection(intervals4, intervals5))
	intervals6 := []Interval{{Start: 6, End: 7}, {Start: 2, End: 4}, {Start: 8, End: 12}}
	fmt.Println(hasOverlap(intervals6))
	intervals7 := []Interval{{Start: 4, End: 5}, {Start: 2, End: 3}, {Start: 3, End: 6}, {Start: 5, End: 7}, {Start: 7, End: 8}}
	fmt.Println(overlapIntervals(intervals7))
	intervals8 := []Interval{{Start: 4, End: 5}, {Start: 2, End: 3}, {Start: 3, End: 5}}
	fmt.Println(minMeetings(intervals8))
	t1 := Interval{Start: 1, End: 4}
	jobs := []job{{Load: 2, Interval: t1}, {Load: 1, Interval: Interval{Start: 2, End: 4}}, {Load: 5, Interval: Interval{Start: 3, End: 6}}}
	fmt.Println(maxLoad(jobs))
	// p1 := [][]Interval{{{Start: 1, End: 3}}}
	// p2 := [][]Interval{{{Start: 9, End: 12}}}
	// p3 := [][]Interval{{{Start: 2, End: 4}, {Start: 6, End: 8}}}
	// h := [][][]Interval{p1, p2, p3}
}
