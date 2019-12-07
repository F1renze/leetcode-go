package q200


func numIslands(grid [][]byte) int {
	if len(grid) < 1 {
		return 0
	}

	count := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == byte('1') {
				count++
				dfs(grid, i, j)
			}
		}
	}
	return count
}

func dfs (grid [][]byte, i, j int) {
	switch {
	case i < 0:
		fallthrough
	case j < 0:
		fallthrough
	case i >= len(grid):
		fallthrough
	case j >= len(grid[i]):
		fallthrough
	case grid[i][j] != byte('1'):
		return
	}
	grid[i][j] = byte('2')
	dfs(grid, i+1, j)
	dfs(grid, i-1, j)
	dfs(grid, i, j+1)
	dfs(grid, i, j-1)
}