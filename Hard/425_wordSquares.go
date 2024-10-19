type Trie struct {
	Children map[byte]*Trie
	Words    []string
}

func NewTrie() *Trie {
	return &Trie{Children: make(map[byte]*Trie)}
}

func (t *Trie) Insert(word string) {
	curr := t

	for i := 0; i < len(word); i++ {
		char := word[i]

		if _, ok := curr.Children[char]; !ok {
			curr.Children[char] = NewTrie()
		}

		curr = curr.Children[char]
		curr.Words = append(curr.Words, word)
	}
}

func (t *Trie) FindByPref(pref string) []string {
	curr := t

	for i := 0; i < len(pref); i++ {
		char := pref[i]

		if _, ok := curr.Children[char]; !ok {
			return []string{}
		}

		curr = curr.Children[char]
	}

	return curr.Words
}

func wordSquares(words []string) [][]string {
	var res [][]string
	wordLen := len(words[0])
	if len(words) == 0 {
		return res
	}

	trie := NewTrie()
	for _, word := range words {
		trie.Insert(word)
	}

	var backtrack func(squareWords []string)
	backtrack = func(squareWords []string) {
		n := len(squareWords)
		if n == wordLen {
			copySquareWords := make([]string, n)
			copy(copySquareWords, squareWords)
			res = append(res, copySquareWords)
			return
		}

		i := n
		var pref string
		for _, word := range squareWords {
			pref += string(word[i])
		}

		for _, word := range trie.FindByPref(pref) {
			squareWords = append(squareWords, word)
			backtrack(squareWords)
			squareWords = squareWords[:len(squareWords)-1]
		}
	}

	for _, word := range words {
		squareWords := []string{word}
		backtrack(squareWords
	}

	return res
}

// m - total number of charrecters
// time: O(m)
// space: O(m)