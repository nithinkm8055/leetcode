package Medium

func searchRange(nums []int, target int) []int {
	return []int{BinSearch(nums, target, false), BinSearch(nums, target, true)}
}

// if max is set, search for last position, else search first position
func BinSearch(nums []int, ele int, max bool) int {

	left := 0
	right := len(nums) - 1
	ans := -1

	for left <= right {
		mid := (left + right) / 2

		if ele == nums[mid] {
			ans = mid
			if max {
				left = mid + 1
			} else {
				right = mid - 1
			}

		} else if ele > nums[mid] {

			left = mid + 1
		} else if ele < nums[mid] {
			right = mid - 1
		}
	}
	return ans
}
