func characterReplacement(s string, k int) int {
	window := make(map[byte]int)
	left := 0
	maxFreq := 0
	res := 0

	for right := 0; right < len(s); right++ {
		window[s[right]]++
		maxFreq = max(maxFreq, window[s[right]])

		if (right-left+1)-maxFreq > k {
			window[s[left]]--
			left++
		}

		res = max(res, right-left+1)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// time: O(N)
// space: O(N)