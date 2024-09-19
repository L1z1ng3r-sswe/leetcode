func solveSudoku(board [][]byte) {
	var backtrack func() bool
	backtrack = func() bool {
		for row := 0; row < 9; row++ {
			for col := 0; col < 9; col++ {
				if board[row][col] == '.' {
					for num := byte('1'); num <= '9'; num++ {
						if isValid(board, row, col, num) {
							board[row][col] = num
							if backtrack() {
								return true
							}
							board[row][col] = '.'
						}
					}

					return false
				}
			}
		}

		return true
	}

	backtrack()
}

func isValid(board [][]byte, row, col int, num byte) bool {
	// check the row
	for i := 0; i < 9; i++ {
		if board[row][i] == num {
			return false
		}
	}

	// check the column
	for i := 0; i < 9; i++ {
		if board[i][col] == num {
			return false
		}
	}

	// check the 3x3 sub-grid
	row = row / 3 * 3
	col = col / 3 * 3

	for i := row; i < row+3; i++ {
		for j := col; j < col+3; j++ {
			if board[i][j] == num {
				return false
			}
		}
	}

	return true
}

// time: O(9^81)
// space: O(1)