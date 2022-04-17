package Hard

func Trap(height []int) int {
	maxHeightLeft := findMaxHeightLeft(height)
	maxHeightRight := findMaxHeightRight(height)

	ans := 0
	for i := 1; i < len(height)-1; i++ {
		ans +=  max(0, min(maxHeightLeft[i], maxHeightRight[i]) - height[i])
	}
	return ans
}

func findMaxHeightLeft(height []int) []int{
	maxHeightLeft := make([]int, len(height))
	maxHeight := 0
	for i := range height {
		maxHeightLeft[i] = maxHeight
		if height[i] > maxHeight {
			maxHeight = height[i]
		}
	}
	return maxHeightLeft
}

func findMaxHeightRight(height []int) []int{
	maxHeightRight := make([]int, len(height))
	maxHeight := 0
	for i := len(height) - 1; i >=0; i -- {
		maxHeightRight[i] = maxHeight
		if height[i] > maxHeight {
			maxHeight = height[i]
		}
	}
	return maxHeightRight
}

func max (a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}



