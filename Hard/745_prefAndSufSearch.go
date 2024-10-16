type WordFilter struct {
	pref *Trie
	suff *Trie
}

func Constructor(words []string) WordFilter {
	wf := WordFilter{
		pref: &Trie{children: make(map[byte]*Trie)},
		suff: &Trie{children: make(map[byte]*Trie)},
	}

	for i, word := range words {
		wf.pref.insert(word, i)
		wf.suff.insert(reverse(word), i)
	}

	return wf
}

func (this *WordFilter) F(pref string, suff string) int {
	prefIndi := this.pref.search(pref)
	if prefIndi == nil {
		return -1
	}

	suffIndi := this.suff.search(reverse(suff))
	if suffIndi == nil {
		return -1
	}

	// prefIndi and suffInde are in non decreasing order -> two pointers technique
	i, j := len(prefIndi)-1, len(suffIndi)-1 // this order because the biggest value is at the end
	for i >= 0 && j >= 0 {
		if prefIndi[i] == suffIndi[j] {
			return prefIndi[i]
		} else if prefIndi[i] > suffIndi[j] {
			i--
		} else {
			j--
		}
	}

	return -1
}

type Trie struct {
	indices  []int
	children map[byte]*Trie
}

func (t *Trie) insert(word string, idx int) {
	curr := t

	for i := 0; i < len(word); i++ {
		ch := word[i]

		if _, ok := curr.children[ch]; !ok {
			curr.children[ch] = &Trie{children: make(map[byte]*Trie)}
		}

		curr = curr.children[ch]
		curr.indices = append(curr.indices, idx)
	}
}

func (t *Trie) search(word string) []int {
	curr := t

	for i := 0; i < len(word); i++ {
		ch := word[i]

		if _, ok := curr.children[ch]; !ok {
			return nil
		}

		curr = curr.children[ch]
	}

	return curr.indices
}

func reverse(word string) string {
	runes := []rune(word)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// constuctor: O(n*L)
// F: O(P+S+k)