package Easy

func IsValid(s string) bool {

	var q []rune
	for _, c := range s {

		switch c {
		case '(':
			q = append(q, ')')
		case '{':
			q = append(q, '}')
		case '[':
			q = append(q, ']')
		default:
			if len(q) == 0 || c != q[len(q)-1] {
				return false
			}
			q = q[:len(q)-1]
		}

	}

	return len(q) == 0
}

