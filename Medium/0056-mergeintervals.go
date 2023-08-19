package Medium

func Merge(intervals [][]int) [][]int {
	//fmt.Println(sort.Ints(intervals))
	result := make([][]int, 0)
	for i := range intervals {
		for j := i + 1; j < len(intervals); j++ {
			temp := make([]int, 0)
			if intervals[i][1] <= intervals[j][1] && intervals[i][1] >= intervals[j][0] {
				temp = append(temp, intervals[i][0])
				temp = append(temp, intervals[j][1])
				result = append(result, temp)
			}
		}
	}
	return result
}
