func smallestRange(lists [][]int) []int {
	maxVal := math.MinInt32

	minHeap := &MinHeap{}
	for i, list := range lists {
		heap.Push(minHeap, &Item{List: i, Idx: 0, Val: list[0]})

		if maxVal < list[0] {
			maxVal = list[0]
		}
	}

	rangeStart, rangeEnd := 0, math.MaxInt32

	for minHeap.Len() > 0 {
		minVal := heap.Pop(minHeap).(*Item)

		if maxVal-minVal.Val < rangeEnd-rangeStart {
			rangeStart = minVal.Val
			rangeEnd = maxVal
		}

		if minVal.Idx+1 < len(lists[minVal.List]) {
			nextVal := &Item{List: minVal.List, Idx: minVal.Idx + 1, Val: lists[minVal.List][minVal.Idx+1]}

			heap.Push(minHeap, nextVal)

			if nextVal.Val > maxVal {
				maxVal = nextVal.Val
			}
		} else {
			break
		}
	}

	return []int{rangeStart, rangeEnd}
}

type Item struct {
	List int
	Idx  int
	Val  int
}

type MinHeap []*Item

func (h MinHeap) Len() int { return len(h) }

func (h MinHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }

func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(*Item))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	last := old[n-1]
	*h = old[:n-1]
	return last
}