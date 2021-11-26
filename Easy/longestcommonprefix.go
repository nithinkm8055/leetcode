package Easy

import "sort"

// Q14 : https://leetcode.com/problems/longest-common-prefix/


// bruteforce : 0ms - faster with more memory
func LongestCommonPrefix(strs []string) string {
	sort.Strings(strs)
	if len(strs) == 1 {
		return strs[0]
	} else if len(strs) == 0 {
		return ""
	}

	highestprefix := strs[0]
	result := ""
	for i := 1 ; i < len(strs) ; i++ {
		for j := 0; j < lowest(strs[i], highestprefix) ; j++ {
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