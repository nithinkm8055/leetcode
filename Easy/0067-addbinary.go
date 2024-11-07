package Easy

import "strconv"

func addBinary(a string, b string) string {
	i := len(a) - 1
	j := len(b) - 1
	carry := 0
	result := ""

	for i >= 0 || j >= 0 || carry == 1 {
		sum := carry
		if i >= 0 {
			sum += int(rune(a[i]) - '0')
			i -= 1
		}
		if j >= 0 {
			sum += int(rune(b[j]) - '0')
			j -= 1
		}

		rem := sum % 2
		result = strconv.Itoa(rem) + result
		carry = sum / 2
	}

	return result
}
