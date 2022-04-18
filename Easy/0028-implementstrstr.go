package Easy

// Q28: https://leetcode.com/problems/implement-strstr/

func StrStr(haystack string, needle string) int {

	if len(needle) == 0 {
		return 0
	}

	for x := 0; x <= len(haystack)-len(needle); x++ {
		if haystack[x:x+len(needle)] == needle {
			return x
		}
	}
	return -1

}
