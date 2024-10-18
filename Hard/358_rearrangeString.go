func rearrangeString(s string, k int) string {
	freqMap := make(map[rune]int)
	for _, char := range s {
		freqMap[char]++
	}

	maxHeap := &MaxHeap{}
	heap.Init(maxHeap)
	for char, freq := range freqMap {
		heap.Push(maxHeap, Item{char: char, freq: freq})
	}

	queue := []Item{}
	var res []rune

	for maxHeap.Len() > 0 {
		mostItem := heap.Pop(maxHeap).(Item)
		res = append(res, mostItem.char)
		mostItem.freq--

		queue = append(queue, mostItem)

		if len(queue) == k {
			front := queue[0]
			queue = queue[1:]
			if front.freq > 0 {
				heap.Push(maxHeap, front)
			}
		}
	}

	if len(res) != len(s) {
		return ""
	}

	return string(res)
}

type Item struct {
	char rune
	freq int
}

type MaxHeap []Item

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

// d - number of distinct elements
// time: O(nlogd)
// space: O(n)