func lengthOfLongestSubstring(s string) int {
	hashTable := make(map[rune]int)
	left := 0
	result := 0

	for right, elem := range s {
		if pos, contains := hashTable[elem]; contains && pos >= left {
			left = pos + 1
		}

		hashTable[elem] = right
		result = max(result, right-left+1)
	}

	return result
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

// time: O(S)
// space: O(N)