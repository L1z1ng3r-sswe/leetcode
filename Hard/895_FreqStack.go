type FreqStack struct {
	maxFreq   int
	freqGroup map[int][]int // freq: list of elements
	numFreq   map[int]int   // num: freq
}

func Constructor() FreqStack {
	return FreqStack{
		freqGroup: make(map[int][]int),
		numFreq:   make(map[int]int),
	}
}

func (this *FreqStack) Push(val int) {
	this.numFreq[val]++
	freq := this.numFreq[val]
	this.freqGroup[freq] = append(this.freqGroup[freq], val)
	this.maxFreq = max(this.maxFreq, freq)
}

func (this *FreqStack) Pop() int {
	maxFreqElem := this.freqGroup[this.maxFreq][len(this.freqGroup[this.maxFreq])-1]

	// Remove the element from the current max frequency group
	this.freqGroup[this.maxFreq] = this.freqGroup[this.maxFreq][:len(this.freqGroup[this.maxFreq])-1]

	// Decrement its frequency in numFreq map
	this.numFreq[maxFreqElem]--

	// If no elements left for the current maxFreq, reduce maxFreq
	if len(this.freqGroup[this.maxFreq]) <= 0 {
		this.maxFreq--
	}

	return maxFreqElem
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// push: O(1)
// pop: O(1)
// space: O(n)