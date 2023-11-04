package Medium

func maxArea(height []int) int {
	maxArea := -1

	i := 0
	j := len(height) - 1

	for i < len(height) && j >= 0 {

		maxArea = max(maxArea, (j-i)*min(height[i], height[j]))

		if height[i] < height[j] {
			i++
		} else if height[i] > height[j] {
			j--
		} else {
			i++
			j--
		}
	}

	return maxArea
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
