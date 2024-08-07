type MinHeap []int

func (mh MinHeap) Len() int           { return len(mh) }
func (mh MinHeap) Less(i, j int) bool { return mh[i] < mh[j] }
func (mh MinHeap) Swap(i, j int)      { mh[i], mh[j] = mh[j], mh[i] }

func (mh *MinHeap) Push(x interface{}) {
	*mh = append(*mh, x.(int))
}

func (mh *MinHeap) Pop() interface{} {
	old := *mh
	lastI := len(old) - 1
	last := old[lastI]
	*mh = old[:lastI]
	return last
}

func minMeetingRooms(intervals [][]int) int {
	// sort by the start of the meetings (ascending)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// define heap and push there 0 intedex intervals (the first meeting)
	minHeap := &MinHeap{}
	heap.Init(minHeap)

	heap.Push(minHeap, intervals[0][1])

	// check whether the interval start is more than the curr 0 in the minheap, if yeah -> pop from the intervals
	for _, interval := range intervals[1:] {
		if (*minHeap)[0] < interval[0] {
			_ = heap.Pop(minHeap)
		}

		heap.Push(minHeap, interval[1])
	}

	return minHeap.Len()
}

// time: O(NlogN)
// space: O(N)