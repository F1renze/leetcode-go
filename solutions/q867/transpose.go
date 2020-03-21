package q867

/**
  类似于 numpy.T(arr)
 */
func transpose(A [][]int) [][]int {
	y := len(A)
	x := len(A[0])

	T := make([][]int, x)

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			T[i] = append(T[i], A[j][i])
		}
	}
	return T
}
