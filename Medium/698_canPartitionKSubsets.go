func canPartitionKSubsets(nums []int, k int) bool {
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}

	if totalSum%k != 0 {
		return false
	}

	target := totalSum / k
	sort.Sort(sort.Reverse(sort.IntSlice(nums))) // Sort in descending order to optimize
	used := make([]bool, len(nums))

	var dfs func(start int, k int, currSum int) bool
	dfs = func(start int, k int, currSum int) bool {
		if k == 1 {
			return true
		}

		if currSum == target {
			return dfs(0, k-1, 0) // 0 value because the value to get target can be the end of the array, f.g.: (4,1) start = 5
		}

		for i := start; i < len(nums); i++ {
			if used[i] || currSum+nums[i] > target {
				continue
			}

			used[i] = true
			if dfs(i+1, k, currSum+nums[i]) {
				return true
			}
			used[i] = false
		}

		return false
	}

	return dfs(0, k, 0)
}

// time: k*(2**N)
// space: O(N)