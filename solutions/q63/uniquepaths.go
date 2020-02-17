package q63

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])

	path := make([]int, n)
	path[0] = 1

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				path[j] = 0
			} else if j > 0 {
				path[j] += path[j-1]
			}
		}
	}
	return path[n-1]
}
