type PairHeap [][3]int

func (h PairHeap) Len() int           { return len(h) }
func (h PairHeap) Less(i, j int) bool { return h[i][2] < h[j][2] }
func (h PairHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *PairHeap) Push(x interface{}) {
	*h = append(*h, x.([3]int))
}

func (h *PairHeap) Pop() interface{} {
	old := *h
	lastIndex := len(old) - 1
	item := old[lastIndex]
	*h = old[:lastIndex]
	return item
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	var result [][]int
	minHeap := &PairHeap{}
	heap.Init(minHeap)

	for i := 0; i < len(nums1) && i < k; i++ {
		heap.Push(minHeap, [3]int{i, 0, nums1[i] + nums2[0]})
	}

	for minHeap.Len() > 0 && len(result) < k {
		item := heap.Pop(minHeap).([3]int)
		i, j := item[0], item[1]

		result = append(result, []int{nums1[i], nums2[j]})

		if j+1 < len(nums2) {
			heap.Push(minHeap, [3]int{i, j + 1, nums1[i] + nums2[j+1]})
		}
	}

	return result
}

// time: O(kLogk)
// space: O(k)