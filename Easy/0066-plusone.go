package Easy

// Q66: https://leetcode.com/problems/plus-one/submissions/
// 0ms memory and time
func PlusOne(nums []int) []int {

	carry := 0

	for j := len(nums)-1 ; j >= 0; j-- {
		if j == len(nums)-1 {
			nums[j] = nums[j] + 1
		} else {
			nums[j] = nums[j] + carry
		}
		if nums[j] > 9 {
			carry = 1
			nums[j] = 0
		} else {
			carry = 0
		}
	}

	// only append to slice if you have a carry
	if carry == 1 {
		nums = append(nums, 1)
		for i:=0 ;i < len(nums)-1; i++ {
			nums[i+1] = nums[i]
		}
		nums[0] = 1
	}

	return nums

}