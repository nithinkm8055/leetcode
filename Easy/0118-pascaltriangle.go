package Easy

import (
	"fmt"
	"math"
)

func Generate(numRows int) [][]int {

	resultMap := make([][]int, numRows)
	initializeMap(resultMap)
	fmt.Print(math.MaxInt)
	resultMap[0][0] = 1
	for i := 1; i < len(resultMap); i++ {
		for j := 0 ; j <= i; j++ {
			if j == 0 || j == i {
				resultMap[i][j] = 1
			} else {
				resultMap[i][j] = resultMap[i-1][j-1] + resultMap[i-1][j]
			}
		}
	}
	return resultMap

}

func initializeMap(resultMap [][]int) {
	for i := range resultMap {
		resultMap[i] = make([]int, i+1)
	}
}