func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}

	indexOfS := 0

	for i := 0; i < len(t); i++ {
		charS := s[indexOfS]
		charT := t[i]

		if charS == charT {
			indexOfS += 1
		}

		if indexOfS == len(s) {
			return true
		}
	}

	return false
} 