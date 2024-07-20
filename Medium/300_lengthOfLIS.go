func lengthOfLIS(nums []int) int {
	var res = 1

	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}

	for i := range nums {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}

	return res
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

// N = len(nums)
// time: O(N**N)
// space: O(N)