package q62

func uniquePaths(m int, n int) int {
	tmp := make([]int, n)
	for i := range tmp {
		tmp[i] = 1
	}
	// 位置 (i, j) 的路径数 = (i, j-1) + (i-1, j)
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			tmp[j] += tmp[j-1]
		}
	}

	return tmp[n-1]
}
