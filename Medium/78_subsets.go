func subsets(nums []int) [][]int {
	var res [][]int

	var dfs func(currNums1 []int, index int)
	dfs = func(currNums1 []int, index int) {
		currNums := make([]int, len(currNums1))
		copy(currNums, currNums1)
		res = append(res, currNums)

		for i := index; i < len(nums); i++ {
			currNums = append(currNums, nums[i])
			dfs(currNums, i+1)
			currNums = currNums[:len(currNums)-1]
		}
	}

	dfs([]int{}, 0)

	return res
}