package Easy

func SubarraySum(nums []int, k int) int {
	counter := 0
	for i := range nums {
		currSum := 0
		for j := i; j < len(nums); j++ {
			currSum = currSum + nums[j]
			if currSum == k {
				counter++
			}
		}
	}
	return counter
}
