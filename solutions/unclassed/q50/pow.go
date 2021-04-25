package q50

func Pow(x float64, n int) float64 {
	if n < 0 {
		return Pow(1/x, -n)
	}
	return FastPow(x, n)
}

/**
O(log(n)) 快速幂
*/
func FastPow(x float64, n int) float64 {
	if n == 0 {
		return 1.0
	}

	half := FastPow(x, n/2)
	if n%2 == 0 {
		return half * half
	} else {
		return half * half * x
	}
}
