package Easy

import "strconv"

// Q9 : https://leetcode.com/problems/palindrome-number/

// bruteforce : 33ms
func IsPalindrome(x int) bool {
	num := strconv.Itoa(x)
	for i := 0; i < len(num) / 2 ; i++ {
		if num[i] != num[len(num) - 1 - i] {
			return false
		}
	}
	return true
}

func IsPalindrome2(x int) bool {

	if x < 0 || (x % 10 == 0 && x != 0){
		return false
	}

	revertedNumber := 0
	for ;x > revertedNumber; {
		revertedNumber = revertedNumber * 10 + x % 10
		x /= 10
	}

	return x == revertedNumber || x == revertedNumber/10
}