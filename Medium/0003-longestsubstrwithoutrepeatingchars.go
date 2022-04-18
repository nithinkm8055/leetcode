package Medium

import "math"

func LengthOfLongestSubstring(s string) int {
	freq := make(map[rune]int)
	left := 0
	res := 0
	for index, value := range s {
		freq[value] = freq[value] + 1
		for freq[value] > 1 {
			freq[rune(s[left])] = freq[rune(s[left])] - 1
			left++
		}
		res = int(math.Max(float64(res), float64(index-left+1)))
	}

	return res
}
