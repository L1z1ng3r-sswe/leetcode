func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	var intersections [][]int
	i, j := 0, 0

	for i < len(firstList) && j < len(secondList) {
		a1, a2 := firstList[i][0], firstList[i][1]
		b1, b2 := secondList[j][0], secondList[j][1]

		if a1 <= b2 && b1 <= a2 {
			start := max(a1, b1)
			end := min(a2, b2)

			intersections = append(intersections, []int{start, end})
		}

		if a2 > b2 {
			j++
		} else {
			i++
		}
	}

	return intersections
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// time: O(m+n)
// space: O(1)

