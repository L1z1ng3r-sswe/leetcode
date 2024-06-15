func findTargetSumWays(nums []int, target int) int {
	var cache = make(map[string]int)

	var dfs func(index int, total int) int
	dfs = func(index int, total int) int {
		if index == len(nums) {
			if total == target {
				return 1
			}

			return 0
		}

		if res, ok := cache[fmt.Sprintf("%d,%d", index, total)]; ok {
			return res
		}
		temp := dfs(index+1, total-nums[index]) + dfs(index+1, total+nums[index])
		cache[fmt.Sprintf("%d,%d", index, total)] = temp
		return temp
	}
	return dfs(0, 0)
}