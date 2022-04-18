package Easy

// Q53: https://leetcode.com/problems/maximum-subarray/
func MaxSubArray(nums []int) int {
	currSum := 0
	maxSum := nums[0]
	for _, val := range nums {
		currSum += val
		if currSum > maxSum {
			maxSum = currSum
		}
		if currSum < 0 {
			currSum = 0
		}
	}
	return maxSum
}
