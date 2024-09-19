func smallestRange(lists [][]int) []int {
	minHeap := &MinHeap{}

	maxVal := &Item{Val: math.MinInt32}
	for i, list := range lists {
		item := &Item{Val: list[0], List: i, Idx: 0}

		heap.Push(minHeap, item)
		if maxVal.Val < item.Val {
			maxVal = item
		}
	}

	var rangeStart, rangeEnd = 0, math.MaxInt32

	for minHeap.Len() > 0 {
		minVal := heap.Pop(minHeap).(*Item)

		if maxVal.Val-minVal.Val < rangeEnd-rangeStart {
			rangeStart = minVal.Val
			rangeEnd = maxVal.Val
		}

		if minVal.Idx+1 < len(lists[minVal.List]) {
			next := lists[minVal.List][minVal.Idx+1]
			item := &Item{Val: next, List: minVal.List, Idx: minVal.Idx + 1}
			heap.Push(minHeap, item)

			if item.Val > maxVal.Val {
				maxVal = item
			}
		} else {
			break
		}
	}

	return []int{rangeStart, rangeEnd}
}

type Item struct {
	Val  int
	List int
	Idx  int
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

// time: O(nlogk)
// space: O(k)