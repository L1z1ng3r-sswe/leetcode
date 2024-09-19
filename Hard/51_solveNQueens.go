func solveNQueens(n int) [][]string {
	var board [][]byte
	for row := 0; row < n; row++ {
		board = append(board, []byte{})

		for col := 0; col < n; col++ {
			board[row] = append(board[row], '.')
		}
	}

	var res [][]string

	var backtrack func(row int)
	backtrack = func(row int) {
		if row >= n {
			var boardStr []string

			for _, row := range board {
				boardStr = append(boardStr, string(row))
			}
			res = append(res, boardStr)

			return
		}

		for col := 0; col < n; col++ {
			if isValid(board, row, col, n) {
				board[row][col] = 'Q'
				backtrack(row + 1)
				board[row][col] = '.'
			}
		}
	}

	backtrack(0)

	return res
}

func isValid(board [][]byte, row, col, n int) bool {
	// check the column
	for i := 0; i < row; i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}

	// check the leftDiagonal
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 'Q' {
			return false
		}
	}

	// check the rightDiagonal
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if board[i][j] == 'Q' {
			return false
		}
	}

	return true
}

// time: O(n * n!)
// space: O(n! * n^2)


