func construct2DArray(original []int, m int, n int) [][]int {
	var res [][]int //returning array
	var row []int   //row - item of res
	var t int       //counter that is not longer than n - colum

	for _, num := range original {
		row = append(row, num)
		t++

		if t == n {
			res = append(res, row)
			t = 0
			row = []int{}
		}
	}

	if len(res) != m || t != 0 { //comparing the length of res to
		res = [][]int{}
	}

	return res
}