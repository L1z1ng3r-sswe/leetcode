func uniqueLetterString(s string) int {
	var n = len(s)

	idx := make(map[byte][]int) // to track for each character its prev1 and prev2 appearence

	for i := 0; i < len(s); i++ {
		char := s[i]
		idx[char] = []int{-1, -1}
	}

	var res int

	for i := 0; i < len(s); i++ {
		char := s[i]

		prev1 := idx[char][0]
		prev2 := idx[char][1]

		res += (prev1 - prev2) * (i - prev1)

		idx[char] = []int{i, prev1} // new prev1 as i and new prev2 as the old prev1
	}

	// compute for the last positions
	for _, arr := range idx {
		prev1 := arr[0]
		prev2 := arr[1]

		res += (prev1 - prev2) * (n - prev1)
	}

	return res
}

// time: O(n)
// space: O(n)