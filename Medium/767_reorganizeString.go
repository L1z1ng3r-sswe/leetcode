func reorganizeString(s string) string {
	var freqMap = make(map[rune]int)
	for _, char := range s {
		freqMap[char]++
	}

	var maxHeap = &MaxHeap{}
	for char, freq := range freqMap {
		if freq > (len(s)+1)/2 {
			return ""
		}
		heap.Push(maxHeap, Item{char, freq})
	}

	var res []rune
	var prev Item

	for maxHeap.Len() > 0 {
		curr := heap.Pop(maxHeap).(Item)
		res = append(res, curr.char)
		curr.freq--

		if prev.freq > 0 {
			heap.Push(maxHeap, prev)
		}

		prev = curr
	}

	return string(res)
}

type MaxHeap []Item

type Item struct {
	char rune
	freq int
}

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].freq > h[j].freq }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Item))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

// freqmap fill
// fill & check the heap
// combining result

// time: O(nlogn)
// space: O(1) // not more than 26 elements