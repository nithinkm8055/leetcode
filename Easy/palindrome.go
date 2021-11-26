package Easy

import "strconv"

// Q9 : https://leetcode.com/problems/palindrome-number/

// bruteforce : 33ms
func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}  else if x >=10  {
		num := strconv.Itoa(x)
		length := len(num) /2
		for i := 0; i < len(num) / 2 ; i++ {
			if num[i] == num[len(num) - 1 - i] {
				if i == length - 1 {
					return true
				}
			} else {
				return false
			}
		}
	} else {
		return true
	}
	return false
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