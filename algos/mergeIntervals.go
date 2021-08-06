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
}
