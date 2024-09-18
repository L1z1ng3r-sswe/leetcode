func maxUniqueSplit(s string) int {
	var hashMap = make(map[string]bool)

	var backtrack func(idx int) int
	backtrack = func(idx int) int {
		if idx >= len(s) {
			return 0
		}

		var maxUniqueSplits int

		for i := idx + 1; i <= len(s); i++ {
			substr := s[idx:i]

			if !hashMap[substr] {
				hashMap[substr] = true
				maxUniqueSplits = max(maxUniqueSplits, 1+backtrack(i))
				delete(hashMap, substr)
			}
		}
		return maxUniqueSplits
	}

	return backtrack(0)
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}

// time: O(2^n)
// space: O(n)