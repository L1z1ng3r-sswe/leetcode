func exist(board [][]byte, word string) bool {
	rows := len(board)
	columns := len(board[0])
	index := 0

	var recursive func(r, c, index int) bool
	recursive = func(r, c, index int) bool {
		if len(word) == index {
			return true
		}

		if r < 0 || c < 0 || r == rows || c == columns || board[r][c] != word[index] {
			return false
		}

		temp := board[r][c]
		board[r][c] = '#'

		found := recursive(r-1, c, index+1) ||
			recursive(r, c-1, index+1) ||
			recursive(r+1, c, index+1) ||
			recursive(r, c+1, index+1)

		board[r][c] = temp
		return found
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			if board[i][j] == word[index] && recursive(i, j, index) {
				return true
			}
		}
	}

	return false
}
