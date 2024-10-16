func medianSlidingWindow(nums []int, k int) []float64 {
	dualHeap := DualHeap{
		maxHeap: &MaxHeap{},
		minHeap: &MinHeap{},
	}

	for i := 0; i < k; i++ {
		dualHeap.addNum(nums[i], i)
	}
	dualHeap.balanceHeaps()

	var left int
	var res []float64

	res = append(res, dualHeap.median())

	for right := k; right < len(nums); right++ {
		dualHeap.prune(left)
		left++

		dualHeap.addNum(nums[right], right)

		dualHeap.balanceHeaps()

		res = append(res, dualHeap.median())
	}

	return res
}

func (h *DualHeap) addNum(num, idx int) {
	if h.minHeap.Len() == 0 || num >= h.minHeap.Top().Val {
		heap.Push(h.minHeap, Item{Val: num, Idx: idx})
	} else {
		heap.Push(h.maxHeap, Item{Val: num, Idx: idx})
	}

	h.balanceHeaps()
}

func (h *DualHeap) median() float64 {
	if h.maxHeap.Len() == h.minHeap.Len() {
		return (float64(h.maxHeap.Top().Val) + float64(h.minHeap.Top().Val)) / 2
	}
	return float64(h.minHeap.Top().Val)
}

func (h *DualHeap) balanceHeaps() {
	if h.maxHeap.Len() > h.minHeap.Len() {
		heap.Push(h.minHeap, heap.Pop(h.maxHeap))
	} else if h.minHeap.Len() > h.maxHeap.Len()+1 {
		heap.Push(h.maxHeap, heap.Pop(h.minHeap))
	}
}

func (h *DualHeap) prune(idx int) {
	for i, item := range *h.minHeap {
		if item.Idx == idx {
			heap.Remove(h.minHeap, i)
			return
		}
	}

	for i, item := range *h.maxHeap {
		if item.Idx == idx {
			heap.Remove(h.maxHeap, i)
			return
		}
	}
}

type DualHeap struct {
	maxHeap *MaxHeap // for smaller elems
	minHeap *MinHeap // for greater elems
}

type MaxHeap []Item

type Item struct {
	Val int
	Idx int
}

func (h MaxHeap) Len() int { return len(h) }

func (h MaxHeap) Less(i, j int) bool { return h[i].Val > h[j].Val }

func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

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

func (h *MaxHeap) Top() Item {
	return (*h)[0]
}

type MinHeap []Item

func (h MinHeap) Len() int { return len(h) }

func (h MinHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }

func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Item))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

func (h *MinHeap) Top() Item {
	return (*h)[0]
}

// time: O(n(logk))
// space: O(k)