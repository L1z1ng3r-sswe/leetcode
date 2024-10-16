type Interval struct {
	Start int
	End   int
}

func employeeFreeTime(schedule [][]*Interval) []*Interval {
	// Step 1: Combine all the intervals
	var totalSchedule []*Interval

	for _, employee := range schedule {
		for _, interval := range employee {
			totalSchedule = append(totalSchedule, interval)
		}
	}

	// Step 2: Sort all the intervals in ascending order by their start time
	sort.Slice(totalSchedule, func(i, j int) bool {
		return totalSchedule[i].Start < totalSchedule[j].Start
	})

	// Step 3: Find gaps between intervals
	var res []*Interval
	lastEnd := totalSchedule[0].End

	for _, currInterval := range totalSchedule[1:] {
		if currInterval.Start > lastEnd { // There's a gap
			res = append(res, &Interval{Start: lastEnd, End: currInterval.Start})
		}
		lastEnd = max(lastEnd, currInterval.End)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// time: O(nlogn)
// space: O(n)