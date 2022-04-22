package Medium

func findMaxLength(nums []int) int {

	maxlen := 0
	for i := range nums {
		zeroes := 0
		ones := 0
		for j := i; j < len(nums); j++ {
			if nums[j] == 0 {
				zeroes++
			} else {
				ones++
			}
			if ones == zeroes {
				maxlen = max(maxlen, j-i +1)
			}
		}
	}

	return maxlen
}

func  max(a int, b int) int{
	if a > b{
		return a
	}
	return b
}