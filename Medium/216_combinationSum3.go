func combinationSum3(k int, n int) [][]int {
	var res [][]int

	var dfs func(start int, currNums []int, remaining int)
	dfs = func(start int, currNums []int, remaining int) {
		if len(currNums) > k || remaining < 0 {
			return
		}

		if remaining == 0 && len(currNums) == k {
			res = append(res, append([]int(nil), currNums...))
			return
		}

		for i := start; i <= 9; i++ {
			currNums = append(currNums, i)
			dfs(i+1, currNums, remaining-i)
			currNums = currNums[:len(currNums)-1]
		}
	}

	dfs(1, []int{}, n)
	return res
}

