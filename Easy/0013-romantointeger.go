package Easy

// Q13: https://leetcode.com/problems/roman-to-integer/

func RomanToInt(s string) int {
	romans := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
	sum := 0

	for i := 0; i < len(s); i++ {
		if i == len(s)-1 {
			sum += romans[string(s[i])]
			break
		}
		if romans[string(s[i])] > romans[string(s[i+1])] {
			sum += romans[string(s[i])]
		} else if romans[string(s[i])] == romans[string(s[i+1])] {
			sum += romans[string(s[i])]
		} else {
			sum += romans[string(s[i+1])] - romans[string(s[i])]
			i += 1
		}
	}
	return sum
}
