func frequencySort(s string) string {
	freqMap := make(map[rune]int)
	for _, char := range s {
		freqMap[char]++
	}

	var maxHeap = &MaxHeap{}
	for elem, freq := range freqMap {
		heap.Push(maxHeap, Elem{elem, freq})
	}

	var res string
	for maxHeap.Len() != 0 {
		elem := heap.Pop(maxHeap).(Elem)
		for i := 0; i < elem.freq; i++ {
			res += string(elem.elem)
		}
	}

	return res
}

type Elem struct {
	elem rune
	freq int
}

type MaxHeap []Elem

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
	*h = append(*h, x.(Elem))
}

func (h *MaxHeap) Pop() interface{} {
	n := len(*h)
	old := *h
	lastElem := old[n-1]
	*h = old[:n-1]
	return lastElem
}

// time: O(nLogn)
// space: O(n)