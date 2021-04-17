package solutions

/**
202. 快乐数
*/

func isHappy(n int) bool {
	m := map[int]bool{}
	sum := 0
	for n != 1 {
		if m[n] {
			return false
		}
		m[n] = true

		for n> 0 {
			rem := n%10
			sum += rem *rem
			n/=10
		}
		n = sum
		sum=0
	}

	return true
}