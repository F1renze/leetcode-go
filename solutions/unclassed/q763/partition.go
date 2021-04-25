package q763

func partitionLabels(S string) []int {
	var (
		s, e int
		r    []int
	)

	if len(S) < 1 {
		return r
	}
	last := make(map[byte]int)
	for i := range S {
		last[S[i]] = i
	}

	for i := range S {
		e = max(e, last[S[i]])
		if e == i {
			r = append(r, e-s+1)
			s = i + 1
		}
	}
	return r
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
