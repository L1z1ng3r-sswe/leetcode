func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	currEnd := points[0][1]
	arrows := 1

	for _, point := range points[1:] {
		if point[0] > currEnd {
			arrows++
			currEnd = point[1]
		}
	}

	return arrows
}

// time: O(NlogN)
// space: O(1)	