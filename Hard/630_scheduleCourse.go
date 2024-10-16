func scheduleCourse(courses [][]int) int {
	sort.Slice(courses, func(i, j int) bool {
		return courses[i][1] < courses[j][1]
	})

	maxHeap := &MaxHeap{}
	var total int

	for _, course := range courses {
		duration, deadline := course[0], course[1]
		total += duration
		heap.Push(maxHeap, duration)

		if total > deadline {
			total -= heap.Pop(maxHeap).(int)
		}
	}

	return maxHeap.Len()
}

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *MaxHeap) Pop() interface{} {
	old := *h
	lastElem := old[len(old)-1]
	old = old[:len(old)-1]
	*h = old
	return lastElem
}

// time: O(nlogn)
// space: O(n)