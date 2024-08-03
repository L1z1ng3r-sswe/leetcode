type Item struct {
	value int
	row   int
	col   int
}

type IntPriorityQueueImpl struct {
	data []*Item
}

func (pq *IntPriorityQueueImpl) Less(i, j int) bool {
	return pq.data[i].value < pq.data[j].value
}

func (pq *IntPriorityQueueImpl) Len() int {
	return len(pq.data)
}

func (pq *IntPriorityQueueImpl) Swap(i, j int) {
	pq.data[i], pq.data[j] = pq.data[j], pq.data[i]
}

func (pq *IntPriorityQueueImpl) Push(x interface{}) {
	pq.data = append(pq.data, x.(*Item))
}

func (pq *IntPriorityQueueImpl) Pop() interface{} {
	old := pq.data
	n := len(old)
	item := old[n-1]
	pq.data = old[0 : n-1]
	return item
}

func (pq *IntPriorityQueueImpl) Empty() bool {
	return pq.Len() == 0
}

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	minHeap := &IntPriorityQueueImpl{}
	heap.Init(minHeap)

	for r := 0; r < min(k, n); r++ {
		heap.Push(minHeap, &Item{
			value: matrix[r][0],
			row:   r,
			col:   0,
		})
	}

	for i := 0; i < k-1; i++ {
		item := heap.Pop(minHeap).(*Item)
		if item.col+1 < n {
			heap.Push(minHeap, &Item{
				value: matrix[item.row][item.col+1],
				row:   item.row,
				col:   item.col + 1,
			})
		}
	}

	return heap.Pop(minHeap).(*Item).value
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// time: O(kLogk)
// space: O(k)