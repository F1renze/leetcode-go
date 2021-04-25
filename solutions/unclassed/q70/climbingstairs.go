package q70

/**
fibonacci 变种，设置好边界条件复用代码即可
*/
func climbStairsIteration(n int) int {
	a := 1
	b := 1
	for n >= 2 {
		a, b = a+b, a
		n--
	}
	return a
}
