package q695

func maxAreaOfIsland(grid [][]int) int {
	if len(grid) < 1 {
		return 0
	}
	max := 0

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				area := explore(grid, i, j)
				if area > max {
					max = area
				}
			}
		}
	}
	return max
}

func explore(grid [][]int, i, j int) int {
	switch {
	case i < 0:
		fallthrough
	case j < 0:
		fallthrough
	case i >= len(grid):
		fallthrough
	case j >= len(grid[i]):
		fallthrough
	case grid[i][j] == 0:
		return 0
	}

	area := 1
	grid[i][j] = 0
	area += explore(grid, i-1, j)
	area += explore(grid, i+1, j)
	area += explore(grid, i, j-1)
	area += explore(grid, i, j+1)

	return area
}
