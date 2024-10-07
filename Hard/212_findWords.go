type TrieNode struct {
	Children map[byte]*TrieNode
	Word     string
}

func (tn *TrieNode) insert(word string) {
	currNode := tn

	for i := 0; i < len(word); i++ {
		letter := word[i]

		if _, ok := currNode.Children[letter]; !ok {
			currNode.Children[letter] = &TrieNode{Children: make(map[byte]*TrieNode)}
		}

		currNode = currNode.Children[letter]
	}

	currNode.Word = word
}

func findWords(board [][]byte, words []string) []string {
	trieRoot := &TrieNode{Children: make(map[byte]*TrieNode)}

	for _, word := range words {
		trieRoot.insert(word)
	}

	var res []string
	directions := [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}

	var dfs func(i, j int, node *TrieNode)
	dfs = func(i, j int, node *TrieNode) {
		letter := board[i][j]

		currNode, ok := node.Children[letter]
		if !ok {
			return
		}

		if currNode.Word != "" {
			res = append(res, currNode.Word)
			currNode.Word = "" // this word is found, so  we need to hide it from others
		}

		board[i][j] = '#'

		for _, direction := range directions {
			newI := i + direction[0]
			newJ := j + direction[1]

			if newI >= 0 && newJ >= 0 && newI < len(board) && newJ < len(board[0]) && board[newI][newJ] != '#' {
				dfs(newI, newJ, currNode)
			}
		}

		board[i][j] = letter
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			dfs(i, j, trieRoot)
		}
	}

	return res
}

// time: O(m * n * 4^L) - L is the avg len of the words
// space: O(W*L) - total number of words by the avg len of the words