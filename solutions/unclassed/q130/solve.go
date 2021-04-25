package q130

func solve(board [][]byte) {
	if len(board) < 1 {
		return
	}

	for i := range board {
		for j := range board[i] {
			isEdge := i == 0 || j == 0 || i == len(board)-1 || j == len(board[i])-1

			if isEdge && board[i][j] == byte('O') {
				dfs(board, i, j)
			}
		}
	}

	for i := range board {
		for j := range board[i] {
			val := board[i][j]

			switch val {
			case byte('O'):
				board[i][j] = byte('X')
			case byte('#'):
				board[i][j] = byte('O')
			}
		}
	}
}

func dfs(board [][]byte, i, j int) {
	switch {
	case i < 0:
		fallthrough
	case j < 0:
		fallthrough
	case i >= len(board):
		fallthrough
	case j >= len(board[i]):
		fallthrough
	case board[i][j] != byte('O'):
		return
	}
	board[i][j] = byte('#')
	dfs(board, i+1, j)
	dfs(board, i-1, j)
	dfs(board, i, j+1)
	dfs(board, i, j-1)
}
