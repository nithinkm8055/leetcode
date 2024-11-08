package Easy

func majorityElement(nums []int) int {
	// Boyer Moores Algorithm
	// only works when it is guaranteed that a majority element exists

	candidate := nums[0]
	count := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != candidate {
			count--
			if count == 0 {
				candidate = nums[i]
				count++
			}
		} else {
			count++
		}
	}

	return candidate
}
