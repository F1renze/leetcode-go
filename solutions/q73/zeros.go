package q73


func setZeroes(matrix [][]int)  {
	m := len(matrix)
	n := len(matrix[0])

	rowMarkQ := make(map[int]struct{}, m)
	colMarkQ := make(map[int]struct{}, n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				rowMarkQ[i] = struct{}{}
				colMarkQ[j] = struct{}{}
			}
		}
	}

	for r, _ := range rowMarkQ {
		flush(matrix, r, true)
	}
	for c, _ := range colMarkQ {
		flush(matrix, c, false)
	}
}

func flush(matrix [][]int, pos int, isRow bool) {
	if isRow {
		for i := range matrix[0] {
			matrix[pos][i] = 0
		}
		return
	}
	for i := range matrix {
		matrix[i][pos] = 0
	}
}
