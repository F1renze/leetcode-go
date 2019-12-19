package q64


func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	for i := m-1; i >= 0; i-- {
		for j := n-1; j >= 0; j-- {
			if i == (m -1) && j == (n-1) {
				continue
			}
			if i == (m-1) {
				grid[i][j] += grid[i][j+1]
			} else if j == (n-1) {
				grid[i][j] += grid[i+1][j]
			} else {
				grid[i][j] += min(grid[i+1][j], grid[i][j+1])
			}
		}
	}
	return grid[0][0]
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
