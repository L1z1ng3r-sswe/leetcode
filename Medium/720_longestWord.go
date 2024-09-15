func longestWord(words []string) string {
	var res string
	sort.Strings(words)

	hashMap := make(map[string]bool)
	hashMap[""] = true

	for _, word := range words {
		if !hashMap[word[:len(word)-1]] {
			continue
		}

		hashMap[word] = true
		if len(word) > len(res) {
			res = word
		}
	}

	return res
}

// time: O(nlogn)
// space: O(n)