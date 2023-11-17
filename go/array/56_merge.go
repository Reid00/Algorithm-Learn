package array

import "sort"

// merge 合并区间
func merge_2(intervals [][]int) [][]int {

	// 排序intervals
	sort.Sort(sortedArray(intervals))

	length := len(intervals)

	for i := 0; i < length-1; {
		cur, next := intervals[i], intervals[i+1]

		if cur[1] >= next[0] {
			if cur[1] < next[1] {
				intervals[i][1] = intervals[i+1][1]
			}

			intervals = append(intervals[:i+1], intervals[i+2:]...)
			length--
			continue
		}
		i++
	}

	return intervals

}

type sortedArray [][]int

func (s sortedArray) Len() int {
	return len(s)
}

func (s sortedArray) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortedArray) Less(i, j int) bool {
	if s[i][0] == s[j][0] {
		return s[i][1] <= s[j][1]
	}
	return s[i][0] <= s[j][0]
}
