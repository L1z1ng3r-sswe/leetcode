func pacificAtlantic(heights [][]int) [][]int {
	m, n := len(heights), len(heights[0])

	pacific := make([][]bool, m)
	atlantic := make([][]bool, m)

	for i := 0; i < m; i++ {
		pacific[i] = make([]bool, n)
		atlantic[i] = make([]bool, n)
	}

	var dfs func(r, c int, ocean [][]bool)
	dfs = func(r, c int, ocean [][]bool) {
		if ocean[r][c] {
			return
		}

		ocean[r][c] = true

		directions := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

		for _, d := range directions {
			nr, nc := r+d[0], c+d[1]

			if nr >= 0 && nr < m && nc >= 0 && nc < n && heights[nr][nc] >= heights[r][c] {
				dfs(nr, nc, ocean)
			}
		}
	}

	for i := 0; i < m; i++ {
		dfs(i, 0, pacific)
		dfs(i, n-1, atlantic)
	}

	for i := 0; i < n; i++ {
		dfs(0, i, pacific)
		dfs(m-1, i, atlantic)
	}

	var res [][]int
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			if pacific[r][c] && atlantic[r][c] {
				res = append(res, []int{r, c})
			}
		}
	}

	return res
}

// define 2 empty heights matrix with false value representing atlantic and pacific
// row, column, ocean
// // check if already visited
// // mark as readed if not visited
// // 4 directions
// // iterate 4 times in each direction
// // // shiftedPosition = currentPosition + shift
// // // shiftedPosition is more than or equal to passed value ? continue : the next direction

// vary within left and right faces
// // row, column = i, first column, pacific
// // row, column = i, last column, atlantic

// vary within top and bottom faces
// // row, column = first row, i, pacific
// // row, column = last row, i, atlantic

// iterate through both oceans and determine whether cells are true in both

// time: O(N*M)
// space: O(N*M)
