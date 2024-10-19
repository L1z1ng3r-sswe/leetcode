type Trie struct {
	Children map[rune]*Trie
	FreqMap  map[string]int
}

func NewTrie() *Trie {
	return &Trie{
		Children: make(map[rune]*Trie),
		FreqMap:  make(map[string]int),
	}
}

func (t *Trie) Insert(word string, freq int) {
	curr := t
	for _, char := range word {
		if _, ok := curr.Children[char]; !ok {
			curr.Children[char] = NewTrie()
		}
		curr = curr.Children[char]
		curr.FreqMap[word] += freq
	}
}

type AutocompleteSystem struct {
	root     *Trie
	currNode *Trie
	currWord strings.Builder
}

func Constructor(sentences []string, times []int) AutocompleteSystem {
	as := AutocompleteSystem{
		root: NewTrie(),
	}
	as.currNode = as.root

	for i, word := range sentences {
		as.root.Insert(word, times[i])
	}

	return as
}

func (this *AutocompleteSystem) Input(c byte) []string {
	if c == '#' {
		this.root.Insert(this.currWord.String(), 1)
		this.currNode = this.root
		this.currWord.Reset()
		return []string{}
	}

	this.currWord.WriteByte(c)

	char := rune(c)

	if _, ok := this.currNode.Children[char]; !ok {
		this.currNode.Children[char] = NewTrie()
	}
	this.currNode = this.currNode.Children[char]

	return findMostFreq(this.currNode.FreqMap)
}

func findMostFreq(freqMap map[string]int) []string {
	freqList := make([]string, 0, len(freqMap))
	for sentence := range freqMap {
		freqList = append(freqList, sentence)
	}

	if len(freqList) <= 3 {
		return freqList
	}

	sort.Slice(freqList, func(i, j int) bool {
		if freqMap[freqList[i]] == freqMap[freqList[j]] {
			return freqList[i] < freqList[j]
		}
		return freqMap[freqList[i]] > freqMap[freqList[j]]
	})

	return freqList[:3]
}

// K - number of sentences in the node
// L - length of the input
// N - nubmer of characters across all the inputs & constuctors
// M - number of unique sentences

// Input: O(L + K log K)
// space: O(N + M * L)

git add .
git commit -m "feat: 642_autocompleteSystem done;  Input: O(L + K log K),  space: O(N + M * L) where K - number of sentences in the node
L - length of the input
N - nubmer of characters across all the inputs & constuctors
M - number of unique sentences"
git push origin main