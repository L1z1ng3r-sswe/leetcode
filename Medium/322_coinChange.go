func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)

	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0

	for _, coin := range coins{
		for i := coin; i <= amount; i++ {
			if i - coin >= 0 {
				dp[i] = min(dp[i-coin]+1, dp[i])
			}
		}
	}

	if dp[amount] == math.MaxInt32 {
		return -1
	}

	return dp[amount]
}

func min(n1, n2 int) int  {
	if n1 < n2 {
		return n1
	}

	return n2
}

