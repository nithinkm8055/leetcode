package Medium

import "sort"

func threeSum(nums []int) [][]int {

	intMap := make(map[int]int)
	ans := make([][]int, 0)

	for i := range nums {
		intMap[nums[i]] = i
	}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			c := -(nums[i] + nums[j])

			if _, ok := intMap[c]; ok && intMap[c] > j {

				arr := []int{nums[i], nums[j], c}
				sort.Ints(arr)

				if !checkDuplicates(ans, arr) {
					ans = append(ans, arr)
				}
			}
		}
	}
	return ans
}

func checkDuplicates(ans [][]int, arr []int) bool {
	for i := range ans {
		if (ans[i][0] == arr[0]) && (ans[i][1] == arr[1]) && (ans[i][2] == arr[2]) {
			return true
		}
	}
	return false
}
