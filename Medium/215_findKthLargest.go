type MaxHeap []int

func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	lastElem := old[n-1]
	*h = old[:n-1]
	return lastElem
}

func findKthLargest(nums []int, k int) int {
	var maxHeap = &MaxHeap{}

	for _, num := range nums {
		heap.Push(maxHeap, num)
	}

	for i := k; i > 1; i-- {
		heap.Pop(maxHeap)
	}

	return heap.Pop(maxHeap).(int)
}

// space: O(N)
// time: O(N)

git add .
git commit -m "feat: 215_findKthLargest done; space: O(N), time: O(N)"
git push origin main
