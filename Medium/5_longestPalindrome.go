func longestPalindrome(s string) string {
	var start, end = 0, 0

	var isPalindrome = func(s string, left, right int) int {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++
		}

		return right - left - 1
	}

	for i := 0; i < len(s); i++ {
		singleLen := isPalindrome(s, i, i)
		doubleLen := isPalindrome(s, i, i+1)
		maxLen := max(singleLen, doubleLen)

		if end-start+1 < maxLen {
			start = i - (maxLen-1)/2
			end = i + maxLen/2
		}
	}

	return s[start : end+1]
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}