package q1128


func numEquivDominoPairs(dominoes [][]int) int {
	m := make(map[int]int)
	r := 0

	for i := range dominoes {
		if dominoes[i][0] < dominoes[i][1] {
			r = dominoes[i][1] * 10 + dominoes[i][0]
		} else {
			r = dominoes[i][0] * 10 + dominoes[i][1]
		}
		m[r]++
	}

	r = 0
	for _, v := range m {
		r += v * (v-1) / 2
	}
	return r
}

