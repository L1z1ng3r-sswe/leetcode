func spiralOrder(matrix [][]int) []int {
	if len(matrix[0]) == 0 {
		return []int{}
	}

	m, n := len(matrix)-1, len(matrix[0])-1
	res := []int{}

	top, bottom := 0, m
	left, right := 0, n

	for left <= right && top <= bottom {
		// to right
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
		top++

		// to down
		if top > bottom {
			break
		}
		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		right--

		// to left
		if left > right {
			break
		}
		for i := right; i >= left; i-- {
			res = append(res, matrix[bottom][i])
		}
		bottom--

		// to top
		if top > bottom {
			break
		}
		for i := bottom; i >= top; i-- {
			res = append(res, matrix[i][left])
		}
		left++
	}

	return res
}
