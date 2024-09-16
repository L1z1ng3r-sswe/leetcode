type Trie struct {
	Next [2]*Trie
}

func (t *Trie) Insert(num int) {
	currTrie := t

	for i := 31; i >= 0; i-- {
		bit := (num >> i) & 1 // 1 or 0
		if currTrie.Next[bit] == nil {
			currTrie.Next[bit] = &Trie{}
		}
		currTrie = currTrie.Next[bit]
	}
}

func (t *Trie) FindMaxXOR(num int) int {
	var res int
	currTrie := t

	for i := 31; i >= 0; i-- {
		bit := (num >> i) & 1 // 1 or 0

		opppsiteBit := 1 - bit
		if currTrie.Next[opppsiteBit] != nil {
			res = (res << 1) | 1 // add 1 to the end
			currTrie = currTrie.Next[opppsiteBit]
		} else {
			res = (res << 1)
			currTrie = currTrie.Next[bit]
		}
	}

	return res
}

func findMaximumXOR(nums []int) int {
	var res int
	trie := &Trie{}

	trie.Insert(nums[0])

	for _, num := range nums[1:] {
		res = max(res, trie.FindMaxXOR(num))
		trie.Insert(num)
	}

	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// time: O(n)
// space: O(n)