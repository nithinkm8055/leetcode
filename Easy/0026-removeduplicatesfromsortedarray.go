package Easy

// Q26: https://leetcode.com/problems/remove-duplicates-from-sorted-array/

// {0,0,1,1,1,2,2,3,3,4}

func RemoveDuplicates(nums []int) int {
	l := 0
	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i] != nums[i-1] {
			nums[l] = nums[i]
			l++
		}
	}
	return l
}
