func frequencySort(s string) string {
	freqMap := make(map[rune]int)
	for _, char := range s {
		freqMap[char]++
	}

	var maxHeap = &MaxHeap{}
	for char, freq := range freqMap {
		heap.Push(maxHeap, Item{char, freq})
	}

	var res string
	for maxHeap.Len() != 0 {
		item := heap.Pop(maxHeap).(Item)
		for i := 0; i < item.freq; i++ {
			res += string(item.char)
		}
	}

	return res
}

type Item struct {
	char rune
	freq int
}

type MaxHeap []Item

func (h MaxHeap) Less(i, j int) bool {
	return h[i].freq > h[j].freq
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h MaxHeap) Len() int {
	return len(h)
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Item))
}

func (h *MaxHeap) Pop() interface{} {
	n := len(*h)
	old := *h
	lastElem := old[n-1]
	*h = old[:n-1]
	return lastElem
}

// time: O(nLogn)
// space: O(1) not more than 26 elements