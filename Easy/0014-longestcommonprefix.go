package Easy

import "sort"

func longestCommonPrefixSol(strs []string) string {
	var result = ""
	for i := 0; ; i++ {
		for _, str := range strs {
			if i == len(str) || strs[0][i] != str[i] {
				return result
			}
		}
		result += string(strs[0][i])
	}
}

// Q14 : https://leetcode.com/problems/longest-common-prefix/
// "flow", "flight", "flower", "dog"
func longestCommonPrefix(strs []string) string {

	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) < len(strs[j])
	})

	result := ""
	for i := 1; i <= len(strs[0]); i++ {
		counter := 0
		for _, str := range strs {
			if len(str) == 0 {
				return result
			}
			if strs[0][0:i] == str[0:i] {
				counter++
			}
		}
		if counter == len(strs) {
			result = strs[0][0:i]
		}
	}
	return result
}

// bruteforce : 0ms - faster with more memory

func LongestCommonPrefix(strs []string) string {
	sort.Slice(strs, func(i, j int) bool {
		return len(strs[i]) < len(strs[j])
	})
	if len(strs) == 1 {
		return strs[0]
	} else if len(strs) == 0 {
		return ""
	}

	highestprefix := strs[0]
	result := ""
	for i := 1; i < len(strs); i++ {
		for j := 0; j < len(strs[0]); j++ {
			if strs[i][j] == highestprefix[j] {
				result += string(strs[i][j])
			} else {
				break
			}
		}
		if result == "" {
			return ""
		}
		highestprefix = result
		result = ""
	}
	return highestprefix
}

func lowest(s1 string, s2 string) int {

	if len(s1) < len(s2) {
		return len(s1)
	}
	return len(s2)
}
