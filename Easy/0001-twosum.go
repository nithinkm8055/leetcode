package Easy

// Q1 : https://leetcode.com/problems/two-sum/

//faster 6ms with more memory usage
func TwoSum2(nums []int, target int) []int {
	result := make(map[int]int)
	for index, num := range nums {
		composite := target - num
		if v, found := result[num]; found {
			return []int{v, index}
		}
		result[composite] = index
	}
	return []int{}
}

var result []int

// slower 32ms with less memory usage -> O(n2)
func TwoSum(nums []int, target int) []int {
	for index, num := range nums {
		check := target - num
		checkComposite(index, check, nums)
		if len(result) > 0 {
			return result
		}
	}

	return result

}

func checkComposite(index int, check int, nums []int) []int {
	for i := index + 1; i < len(nums); i++ {
		if nums[i] == check {
			result = append(result, index, i)
			return result
		}
	}
	return result
}
