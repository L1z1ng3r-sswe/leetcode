func palindromePairs(words []string) [][]int {
	wordMap := make(map[string]int, len(words))

	for i, word := range words {
		wordMap[reverse(word)] = i
	}

	var res [][]int

	for i, word := range words {
		for j := 0; j <= len(word); j++ {
			pref := word[:j] // "",    "l",  "ll", "lls"
			suff := word[j:] // "lls", "ls", "s",  "",
			if isPalindrome(pref) {
				if idx, ok := wordMap[suff]; ok && idx != i {
					res = append(res, []int{idx, i})
				}
			}

			if j != len(word) && isPalindrome(suff) {
				if idx, ok := wordMap[pref]; ok && idx != i {
					res = append(res, []int{i, idx})
				}
			}
		}
	}

	return res
}

func isPalindrome(word string) bool {
	for i, j := 0, len(word)-1; i < j; i, j = i+1, j-1 {
		if word[i] != word[j] {
			return false
		}
	}

	return true
}

func reverse(word string) string {
	runes := []rune(word)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

// time: O(m*n)
// space: O(n)