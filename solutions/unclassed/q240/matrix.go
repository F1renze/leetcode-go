package q240

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) < 1 || len(matrix[0]) < 1 {
		return false
	}
	i, j := len(matrix)-1, 0
	for i >= 0 && j < len(matrix[0]) {
		if matrix[i][j] < target {
			j++
		} else if matrix[i][j] > target {
			i--
		} else {
			return true
		}
	}
	return false
}
