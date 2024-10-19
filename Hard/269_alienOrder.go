
func alienOrder(words []string) string {
	relations := make(map[byte]int)
	graph := make(map[byte][]byte)

	for _, word := range words {
		for i := 0; i < len(word); i++ {
			relations[word[i]] = 0
		}
	}

	for i := 0; i < len(words)-1; i++ {
		word1, word2 := words[i], words[i+1]
		minLen := min(len(word1), len(word2))

		if len(word1) > len(word2) && word1[:minLen] == word2[:minLen] {
			return ""
		}

		for i := 0; i < minLen; i++ {
			parentNode, childNode := word1[i], word2[i]
			if parentNode != childNode {
				graph[parentNode] = append(graph[parentNode], childNode)
				relations[childNode]++
				break
			}
		}
	}

	var queue []byte
	for node, freq := range relations {
		if freq == 0 {
			queue = append(queue, node)
		}
	}

	var res []byte
	for len(queue) != 0 {
		parentNode := queue[0]
		queue = queue[1:]
		res = append(res, parentNode)

		for _, childNode := range graph[parentNode] {
			relations[childNode]--
			if relations[childNode] == 0 {
				queue = append(queue, childNode)
			}
		}
	}

	if len(res) != len(relations) {
		return ""
	}
	return string(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// d - number of characters
// time: O(d)
// space: O(d)