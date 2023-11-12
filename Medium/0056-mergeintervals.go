package Medium

import "sort"

func merge(intervals [][]int) [][]int {

	sort.SliceStable(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	newInterval := intervals[0]

	result := make([][]int, 0)

	for i := range intervals {

		if newInterval[1] < intervals[i][0] {
			result = append(result, newInterval)
			newInterval = intervals[i]
		} else if newInterval[0] > intervals[i][1] {
			result = append(result, intervals[i])
		} else {
			newInterval[0] = min(newInterval[0], intervals[i][0])
			newInterval[1] = max(newInterval[1], intervals[i][1])
		}
	}

	result = append(result, newInterval)

	return result
}
