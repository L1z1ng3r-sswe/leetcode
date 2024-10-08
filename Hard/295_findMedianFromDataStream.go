type MedianFinder struct {
	left  *MaxHeap // for smaller element
	right *MinHeap // for greater elements
}

func Constructor() MedianFinder {
	return MedianFinder{
		left:  &MaxHeap{},
		right: &MinHeap{},
	}
}

func (this *MedianFinder) AddNum(num int) {
	if this.right.Len() == 0 || num >= (*this.right)[0] {
		heap.Push(this.right, num)
	} else {
		heap.Push(this.left, num)
	}

	// balance heaps
	if this.left.Len() > this.right.Len() {
		heap.Push(this.right, heap.Pop(this.left).(int))
	}

	if this.right.Len()-this.left.Len() > 1 {
		heap.Push(this.left, heap.Pop(this.right).(int))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.left.Len() == this.right.Len() {
		return (float64(this.left.MinHeap[0]) + float64((*this.right)[0])) / 2.0
	}

	return float64((*this.right)[0])
}

type MinHeap []int

func (h MinHeap) Len() int { return len(h) }

func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }

func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	lastVal := old[len(old)-1]
	*h = old[:len(old)-1]
	return lastVal
}

type MaxHeap struct {
	MinHeap
}

func (h MaxHeap) Less(i, j int) bool { return h.MinHeap[i] > h.MinHeap[j] }

// insert: O(logn),
// findMedian: time: O(1)