package utils

import "sort"

type Interval [2]int

func (it Interval) Contains(num int) bool {
	return it[0] <= num && num <= it[1]
}

type Intervals []Interval

func (its Intervals) Contains(num int) bool {
	// for _, it := range its {
	// 	if it.Contains(num) {
	// 		return true
	// 	}
	// }
	// return false
	i := sort.Search(
		len(its),
		func(i int) bool { return its[i][1] >= num },
	)
	if i == len(its) {
		return false
	}
	return its[i].Contains(num)
}
