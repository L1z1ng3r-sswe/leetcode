type Point struct {
	distance float64
	index    int
}

type MinHeap []Point

func (h MinHeap) Less(i, j int) bool {
	return h[i].distance < h[j].distance
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h MinHeap) Len() int {
	return len(h)
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Point))
}

func (h *MinHeap) Pop() interface{} {
	n := len(*h)
	old := *h
	lastElem := old[n-1]
	*h = old[:n-1]
	return lastElem
}

func kClosest(points [][]int, k int) [][]int {
	var minHeap = &MinHeap{}

	for i, point := range points {
		hypotenuse := math.Sqrt(float64(point[0]*point[0] + point[1]*point[1]))
		heap.Push(minHeap, Point{hypotenuse, i})
	}

	var res [][]int
	for i := 0; i < k; i++ {
		closestPoint := heap.Pop(minHeap).(Point)
		res = append(res, points[closestPoint.index])
	}

	return res
}

// time: O(nLogn)
// space: O(n)