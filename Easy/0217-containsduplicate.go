package Easy

func containsDuplicate(nums []int) bool {
	eleMap := make(map[int]struct{})
	for i := range nums {
		if _, ok := eleMap[nums[i]]; ok {
			return true
		}
		eleMap[nums[i]] = struct{}{}
	}

	return false
}
