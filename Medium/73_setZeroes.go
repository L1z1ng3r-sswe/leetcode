func setZeroes(matrix [][]int) {
	res := make([][]int, len(matrix)) // [0,0,0,0], [0,0,0,0], [0,0,0,0]

	htr := make(map[int]bool) // [indexRow]bool
	hti := make(map[int]bool) // [indexNum]bool

	for indexRow, row := range matrix { //row - [0,1,2,0]
		res[indexRow] = make([]int, len(row))

		for indexNum, num := range row { //num - 0
			if num == 0 {
				htr[indexRow] = true
				hti[indexNum] = true
			}
		}
	}

	for indexRow, row := range res { //row - [0,0,0,0]

		if htr[indexRow] {
			continue // row - [0,0,0,0]
		}

		copy(row, matrix[indexRow])    // row - [0,1,2,0]
		for indexNum, _ := range row { //num - 0
			if hti[indexNum] {
				row[indexNum] = 0
			}
		}
	}

	copy(matrix, res)
}
