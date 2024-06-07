func subsetsWithDup(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)

	var dfs func(currNums []int, start int)
	dfs = func(currNums []int, start int) {
		// to avoid buggs
		res = append(res, append([]int(nil), currNums...))

		for i := start; i < len(nums); i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}

			// exclude
			currNums = append(currNums, nums[i])
			dfs(currNums, i+1)
			currNums = currNums[:len(currNums)-1]
		}
	}

	dfs([]int{}, 0)

	return res
}
