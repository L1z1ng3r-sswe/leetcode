func findAllConcatenatedWordsInADict(words []string) []string {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	var res []string

	wordSet := make(map[string]bool)
	for _, word := range words {
		if canBeConcatenated(word, wordSet) {
			res = append(res, word)
		}

		wordSet[word] = true
	}

	return res
}

func canBeConcatenated(word string, wordSet map[string]bool) bool {
	dp := make([]bool, len(word)+1)
	dp[0] = true

	for i := 1; i <= len(word); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordSet[word[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(word)]
}

// time: O(nlogn)
// space: O(n)