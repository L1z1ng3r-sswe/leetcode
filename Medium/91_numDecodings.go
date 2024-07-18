func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	dp := make([]int, len(s)+1)
	dp[0] = 1
	dp[1] = 1

	for i := 2; i <= len(s); i++ {
		single := s[i-1 : i]
		double := s[i-2 : i]

		if single != "0" {
			dp[i] += dp[i-1]
		}
		if double >= "10" && double <= "26" {
			dp[i] += dp[i-2]
		}
	}

	return dp[len(s)]
}

// N = len(s)
// time: O(N)
// space: O(N)