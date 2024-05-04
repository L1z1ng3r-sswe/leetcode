//m - number of rows
// n - number of columns
func construct2DArray(original []int, m int, n int) [][]int {
	var result [][]int
	var row []int
	var columnsNumber int // max value is the number of columns (n)

	for _, num := range original {
		row = append(row, num)
		columnsNumber++

		if columnsNumber == n {
			result = append(result, row)
			row = []int{}
			columnsNumber = 0
		}
	}

	if columnsNumber != 0 || m != len(result) {
		result = [][]int{}
	}

	return result
}

func construct2DArray(original []int, columns int, rowlength int) [][]int {
	if len(original) != columns*rowlength {
		return [][]int{}
	}

	var res = make([][]int, columns)
	for i := range res {
		res[i] = make([]int, rowlength)
		copy(res[i], original[i*rowlength:(i+1)*rowlength])
	}

	return res
} 