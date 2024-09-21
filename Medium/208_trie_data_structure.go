type Trie struct {
	Exist bool
	Next  map[byte]*Trie
}

func (t *Trie) Insert(word string) {
	var dfs func(trie *Trie, word string, idx int)
	dfs = func(trie *Trie, word string, idx int) {
		if idx == len(word) {
			trie.Val = word
			return
		}

		if trie.Next == nil {
			trie.Next = make(map[byte]*Trie)
		}

		char := word[idx]

		if _, ok := trie.Next[char]; !ok {
			trie.Next[char] = &Trie{Next: make(map[byte]*Trie)}
		}
		dfs(trie.Next[char], word, idx+1)
	}

	dfs(t, word, 0)
}

func (this *Trie) Search(word string) bool {
	var dfs(trie *Trie, word )
	dfs = func() {
		if idx == len(word) {
			return trie.Val == word
		}

		char := word[i]
		if _, ok := trie.Next[char]; !ok {
			return false
		}

		dfs(trie.Next[char], )
	}

	dfs(t, 0, word)
}

func (this *Trie) StartsWith(prefix string) bool {
	if len(prefix) == 0 {
		return true
	}

	firstByte := prefix[0]
	nextByte, ok := this.Next[firstByte]
	if !ok {
		return false
	}

	return nextByte.StartsWith(prefix[1:])
} 