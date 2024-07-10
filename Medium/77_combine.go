func combine(n int, k int) [][]int {
	var res [][]int

	var dfs func(currNums []int, index int)
	dfs = func(currNums []int, index int) {
		if len(currNums) == k {
			res = append(res, append([]int(nil), currNums...))
			return
		}

		for i := index; i <= n; i++ {
			// add current value
			currNums = append(currNums, i)
			dfs(currNums, i+1)
			// pop
			currNums = currNums[:len(currNums)-1]
		}
	}

	dfs([]int{}, 1)
	return res
}