package Easy

// Q58: https://leetcode.com/problems/length-of-last-word/

func LengthOfLastWord(s string) int {

	count := 0

	for j := len(s) -1 ; j >=0 ; j-- {
		if count > 0 && s[j] == 32 {
			break
		}
		if s[j] != 32 {
			count++
		}
	}
	return count
}