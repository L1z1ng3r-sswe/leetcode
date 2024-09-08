func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int)

	for _, num := range nums {
		freqMap[num]++
	}

	maxHeap := &MaxHeap{}
	for elem, freq := range freqMap {
		heap.Push(maxHeap, [2]int{elem, freq})
	}

	var res []int
	for i := 0; i < k; i++ {
		elem := heap.Pop(maxHeap).([2]int)[0]
		res = append(res, elem)
	}

	return res
}

type MaxHeap [][2]int // elem, freq

func (h MaxHeap) Less(i, j int) bool {
	return h[i][1] > h[j][1]
}

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
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