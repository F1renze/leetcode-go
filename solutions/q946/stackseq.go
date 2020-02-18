package q946

func validateStackSequences(pushed []int, popped []int) bool {
	n := len(pushed)
	if n < 1 {
		return true
	}
	var (
		j int
		s []int
	)

	for i := range pushed {
		s = append(s, pushed[i])

		for len(s) > 0 && j < n && s[len(s)-1] == popped[j] {
			s = s[:len(s)-1]
			j++
		}
	}

	return n == j
}
