func maxSlidingWindow(nums []int, k int) []int {
	var res []int
	maxHeap := &MaxHeap{}

	left, right := 0, 0

	for right < len(nums) {
		heap.Push(maxHeap, &Item{Val: nums[right], Idx: right})

		if right-left+1 >= k {
			for (*maxHeap)[0].Idx < left {
				heap.Pop(maxHeap)
			}

			res = append(res, (*maxHeap)[0].Val)
			left++
		}

		right++
	}

	return res
}

type Item struct {
	Idx int
	Val int
}

type MaxHeap []*Item

func (h MaxHeap) Len() int { return len(h) }

func (h MaxHeap) Less(i, j int) bool { return h[i].Val > h[j].Val }

func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(*Item))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	last := old[n-1]
	*h = old[:n-1]
	return last
}

// time: O(nlogk)
// space: O(n)