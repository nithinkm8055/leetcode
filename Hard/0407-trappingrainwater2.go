package Hard

//████████ ██   ██ ██ ███████      ██████  ██████  ██████  ███████     ██ ███████     ███    ██  ██████  ████████      ██████  ██████  ███    ███ ██████  ██      ███████ ████████ ███████
//   ██    ██   ██ ██ ██          ██      ██    ██ ██   ██ ██          ██ ██          ████   ██ ██    ██    ██        ██      ██    ██ ████  ████ ██   ██ ██      ██         ██    ██
//   ██    ███████ ██ ███████     ██      ██    ██ ██   ██ █████       ██ ███████     ██ ██  ██ ██    ██    ██        ██      ██    ██ ██ ████ ██ ██████  ██      █████      ██    █████
//   ██    ██   ██ ██      ██     ██      ██    ██ ██   ██ ██          ██      ██     ██  ██ ██ ██    ██    ██        ██      ██    ██ ██  ██  ██ ██      ██      ██         ██    ██
//   ██    ██   ██ ██ ███████      ██████  ██████  ██████  ███████     ██ ███████     ██   ████  ██████     ██         ██████  ██████  ██      ██ ██      ███████ ███████    ██    ███████
//
func TrapRainWater(heightMap [][]int) int {

	maxHeightTop := getMaxHeightTop(heightMap)
	maxHeightBottom := getMaxHeightBottom(heightMap)
	maxHeightLeft := getMaxHeightLeft(heightMap)
	maxHeightRight := getMaxHeightRight(heightMap)
	ans := 0

	resultMap := make([][]int, len(heightMap))
	initializeSlice(resultMap, len(heightMap[0]))

	for i := range heightMap {
		for j := range heightMap[i] {
			resultMap[i][j] = max(0, minimum([]int{maxHeightTop[i][j], maxHeightBottom[i][j], maxHeightRight[i][j], maxHeightLeft[i][j]})-heightMap[i][j])
		}
	}

	recursivelyUpdateResultMap(resultMap, heightMap)
	for i := range heightMap {
		for j := range heightMap[i] {

			if resultMap[i][j] > 0 {
				surroundMinimum := minimum([]int{heightMap[i-1][j], heightMap[i+1][j], heightMap[i][j-1], heightMap[i][j+1]})
				if resultMap[i][j] > surroundMinimum {
					resultMap[i][j] = surroundMinimum - heightMap[i][j]
				}
				ans += resultMap[i][j]
			}
		}
	}
	return ans
}


func recursivelyUpdateResultMap(resultMap [][]int, heightMap [][]int) int {
	count := 0

	for i := range heightMap {
		for j := range heightMap[i] {
			if resultMap[i][j] > 0 {

				for k, v := range []int{resultMap[i-1][j], resultMap[i+1][j], resultMap[i][j-1], resultMap[i][j+1]} {
					a := 0
					b := 0

					if k == 0 {
						a = i - 1
						b = j
					} else if k == 1 {
						a = i + 1
						b = j
					} else if k == 2 {
						a = i
						b = j - 1
					} else {
						a = i
						b = j + 1
					}

					if v > 0 {
						if resultMap[i][j] + heightMap[i][j] > v + heightMap[a][b] {
							count ++
							resultMap[i][j] = v + heightMap[a][b] - heightMap[i][j]
						}
					}
				}

			}
		}
	}
	if count == 0{
		return 0
	}
	return recursivelyUpdateResultMap(resultMap, heightMap)
}

func getMaxHeightTop(heightMap [][]int) [][]int {
	maxHeightTop := make([][]int, len(heightMap))
	initializeSlice(maxHeightTop, len(heightMap[0]))
	for j := 0; j < len(heightMap[0]); j++ {
		maxHeight := 0
		for i := 0; i < len(heightMap); i++ {
			maxHeightTop[i][j] = maxHeight
			if heightMap[i][j] > maxHeight {
				maxHeight = heightMap[i][j]
			}
		}
	}
	return maxHeightTop
}

func getMaxHeightBottom(heightMap [][]int) [][]int {
	maxHeightBottom := make([][]int, len(heightMap))
	initializeSlice(maxHeightBottom, len(heightMap[0]))
	for j := 0; j < len(heightMap[0]); j++ {
		maxHeight := 0
		for i := len(heightMap) - 1; i >= 0; i-- {
			maxHeightBottom[i][j] = maxHeight
			if heightMap[i][j] > maxHeight {
				maxHeight = heightMap[i][j]
			}
		}
	}

	return maxHeightBottom
}

func getMaxHeightRight(heightMap [][]int) [][]int {
	maxHeightRight := make([][]int, len(heightMap))
	initializeSlice(maxHeightRight, len(heightMap[0]))
	for i := range heightMap {
		maxHeight := 0
		for j := len(heightMap[0]) - 1; j >= 0; j-- {
			maxHeightRight[i][j] = maxHeight
			if heightMap[i][j] > maxHeight {
				maxHeight = heightMap[i][j]
			}
		}
	}
	return maxHeightRight
}

func getMaxHeightLeft(heightMap [][]int) [][]int {
	maxHeightLeft := make([][]int, len(heightMap))
	initializeSlice(maxHeightLeft, len(heightMap[0]))
	for i := range heightMap {
		maxHeight := 0
		for j := range heightMap[i] {
			maxHeightLeft[i][j] = maxHeight
			if heightMap[i][j] > maxHeight {
				maxHeight = heightMap[i][j]
			}
		}
	}
	return maxHeightLeft
}

func initializeSlice(inputSlice [][]int, colLength int) [][]int {

	for i := range inputSlice {
		inputSlice[i] = make([]int, colLength)
	}
	return inputSlice
}

func minimum(inputSlice []int) int {
	minimum := 30000
	for i := range inputSlice {
		if inputSlice[i] < minimum {
			minimum = inputSlice[i]
		}
	}
	return minimum
}
