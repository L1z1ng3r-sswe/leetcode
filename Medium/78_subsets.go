func subsets(nums []int) [][]int {
	var res [][]int

	var dfs func(currNums []int, index int)
	dfs = func(currNums []int, index int) {
		if index == len(nums) {
			res = append(res, append([]int(nil), currNums...))
			return
		}

		dfs(append(currNums, nums[index]), index+1)
		dfs(currNums, index+1)
	}

	dfs([]int{}, 0)

	return res
}