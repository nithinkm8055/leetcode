package Easy

// Q35: https://leetcode.com/problems/search-insert-position/
func searchInsert(nums []int, target int) int {
	i := 0
	for ; i < len(nums); i++ {
		if nums[i] >= target {
			return i
		}
	}

	return i

}
