package q120

func minimumTotal(triangle [][]int) int {
	if len(triangle) < 1 {
		return 0
	}

	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			triangle[i][j] += min(triangle[i+1][j], triangle[i+1][j+1])
		}
	}
	return triangle[0][0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
