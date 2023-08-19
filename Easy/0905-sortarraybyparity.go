package Easy

func SortArrayByParity(nums []int) []int {

	i := 0

	for j := len(nums) - 1; j > i ; {

		if nums[i] % 2 != 0 {
			swap(nums,i ,j)
			j--
		} else{
			i++
		}

	}

	return nums

}

func swap(nums []int, a , b int) {
	t := nums[a]
	nums[a] = nums[b]
	nums[b] = t

}