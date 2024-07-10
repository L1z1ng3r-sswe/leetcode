func rotate(matrix [][]int) {
	l, r := 0, len(matrix)-1

	for i := 0; i < len(matrix)/2; i++ {
		for j := 0; j < r-l; j++ {
			t, b := l, r

			// temporaray variable for top-left
			topLeft := matrix[t][l+j]

			// move bottom-left to top-left
			matrix[t][l+j] = matrix[b-j][l]

			// move bottom-right to bottom-left
			matrix[b-j][l] = matrix[b][r-j]

			// move top-right to bottom-right
			matrix[b][r-j] = matrix[t+j][r]

			// move temporary variable to top-right
			matrix[t+j][r] = topLeft
		}
		l++
		r--
	}
}
