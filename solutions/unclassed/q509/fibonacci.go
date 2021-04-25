package q509

func Fib(N int) int {

	cache := make(map[int]int)

	return calc(N, cache)
}

func calc(n int, m map[int]int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		if t := m[n]; t != 0 {
			return t
		}
		m[n] = calc(n-1, m) + calc(n-2, m)
		return m[n]
	}
}

func FibIteration(N int) int {
	a := 0
	b := 1
	for N >= 1 {
		a, b = a+b, a
		N--
	}
	return a
}
