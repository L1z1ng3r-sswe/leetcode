func combinationSum(candidates []int, target int) [][]int {
	var res [][]int

	var dfs func(index int, currNums []int, total int)
	dfs = func(index int, currNums []int, total int) {
		if index == len(candidates) || total > target {
			return
		}
		if total == target {
			res = append(res, append([]int(nil), currNums...))
			return
		}

		// append the same currNums
		currNums = append(currNums, candidates[index])
		dfs(index, currNums, total+candidates[index])
		// pop
		currNums = currNums[:len(currNums)-1]
		dfs(index+1, currNums, total)
	}

	dfs(0, []int{}, 0)

	return res
}