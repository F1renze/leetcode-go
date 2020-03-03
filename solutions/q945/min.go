package q945


func minIncrementForUnique(A []int) int {
	m := make([]int, 40000)

	var cnt, max int

	for _, v := range A {
		m[v]++
		if v > max {
			max = v
		}
	}

	for i := 0; i < max; i++ {
		if m[i] > 1 {
			m[i+1] += m[i] - 1
			cnt += m[i] - 1
		}
	}
	if m[max] > 1 {
		a := m[max] - 1
		// Arithmetic progression
		cnt += a * (a + 1) / 2
	}
	return cnt
}
