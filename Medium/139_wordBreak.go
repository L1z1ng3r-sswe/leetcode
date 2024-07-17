func wordBreak(s string, wordDict []string) bool {
	wordTable := make(map[string]bool)
	for _, word := range wordDict {
		wordTable[word] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordTable[s[j:i]] {
				dp[i] = true
			}
		}
	}

	return dp[len(s)]
}

// N = len(s)
// M = len(wordDict)
// time = O(N**N)
// space = O(M+N)