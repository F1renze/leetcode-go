package q174

func calculateMinimumHP(dungeon [][]int) int {
	m := len(dungeon)
	if m < 1 {
		return 0
	}
	n := len(dungeon[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == m-1 && j == n-1 {
				dp[i][j] = max(1, 1-dungeon[i][j])
			} else if i == m-1 {
				dp[i][j] = max(1, dp[i][j+1]-dungeon[i][j])
			} else if j == n-1 {
				dp[i][j] = max(1, dp[i+1][j]-dungeon[i][j])
			} else {
				dp[i][j] = max(1, min(dp[i+1][j], dp[i][j+1])-dungeon[i][j])
			}
		}
	}
	return dp[0][0]
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
