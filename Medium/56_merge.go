func merge(intervals [][]int) [][]int {
	var result [][]int

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	prevInterval := intervals[0]
	for i := 1; i < len(intervals); i++ {
		interval := intervals[i]

		if interval[0] <= prevInterval[1] {
			if interval[1] > prevInterval[1] {
				prevInterval = []int{prevInterval[0], interval[1]}
			}
			continue
		}

		result = append(result, prevInterval)
		prevInterval = interval
	}

	result = append(result, prevInterval)
	return result
}

// time: O(nlogn)
// space: O(n)