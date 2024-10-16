func findSubstring(s string, words []string) []int {
	var res []int
	wordLen := len(words[0])

	wordMap := make(map[string]int)
	for _, word := range words {
		wordMap[word]++
	}

	for i := 0; i < wordLen; i++ {
		right, left := i, i
		window := make(map[string]int)
		var count int

		for right+wordLen <= len(s) {
			word := s[right : right+wordLen]
			right += wordLen

			if _, ok := wordMap[word]; ok {
				count++
				window[word]++

				for window[word] > wordMap[word] {
					leftWord := s[left : left+wordLen]
					window[leftWord]--
					left += wordLen
					count--
				}

				if count == len(words) {
					res = append(res, left)
				}
			} else {
				left = right
				window = make(map[string]int)
				count = 0
			}
		}
	}

	return res
}

// n - len(s)
// k - len(words[0])
// time: O(k*n)
// space: O(n)