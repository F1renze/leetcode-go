package q1021


func removeOuterParentheses(S string) string {
	n := len(S)
	if n < 2 {
		return S
	}

	var (
		buf []byte
		c int
	)

	for i := 0; i < n; i++ {
		if S[i] == '(' {
			c++
			if c > 1 {
				buf = append(buf, S[i])
			}
		} else if S[i] == ')' {
			c--
			if c > 0 {
				buf = append(buf, S[i])
			}
		}

	}
	return string(buf)
}