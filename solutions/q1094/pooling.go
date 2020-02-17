package q1094

func carPooling(trips [][]int, capacity int) bool {
	m := [1001]int{}
	for i := range trips {
		// pick up
		m[trips[i][1]] += trips[i][0]
		// drop off
		m[trips[i][2]] -= trips[i][0]
	}

	for i := 1; i < len(m); i++ {
		m[i] += m[i-1]
		if m[i] > capacity {
			return false
		}
	}

	return true
}
