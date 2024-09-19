func mergeKLists(lists []*ListNode) *ListNode {
	var minHeap = &MinHeap{}

	for _, list := range lists {
		if list != nil {
			heap.Push(minHeap, list)
		}
	}

	dummy := &ListNode{}
	curr := dummy

	for minHeap.Len() > 0 {
		minElem := heap.Pop(minHeap).(*ListNode)

		curr.Next = minElem
		curr = curr.Next

		if minElem.Next != nil {
			heap.Push(minHeap, minElem.Next)
		}
	}

	return dummy.Next
}

type MinHeap []*ListNode

func (h MinHeap) Len() int { return len(h) }

func (h MinHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }

func (h MinHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	last := old[n-1]
	*h = old[:n-1]
	return last
}

// time: O(nlogn)
// space: O(n)