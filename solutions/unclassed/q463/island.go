package q463


func islandPerimeter(grid [][]int) int {
	r, w, h := 0, len(grid), len(grid[0])
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				continue
			}

			if i == 0 || grid[i-1][j] == 0 {
				r += 1
			}
			if i == w-1 || grid[i+1][j] == 0 {
				r += 1
			}
			if j == 0 || grid[i][j-1] == 0 {
				r += 1
			}
			if j == h-1 || grid[i][j+1] == 0 {
				r += 1
			}

		}
	}

	return r
}


