func canPartition(nums []int) bool {
	var total int
	for _, num := range nums {
		total += num
	}

	if total%2 != 0 {
		return false
	}

	target := total / 2

	dp := make([]bool, target+1)
	dp[0] = true

	for _, num := range nums {
		for remaining := target; remaining >= num; remaining-- {
			dp[remaining] = dp[remaining] || dp[remaining-num]
		}
	}

	return dp[target]
}

// N = len(nums)
// M = N/2
// time = O(N*M)
// space = O(M)