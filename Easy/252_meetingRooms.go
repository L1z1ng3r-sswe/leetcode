func CanAttendMeetings(intervals []*Interval) bool {
	sorted := bubbleSort(intervals)

	for i := 0; i < len(sorted)-1; i++ {
		if sorted[i].End >= sorted[i+1].Start {
			return false
		}
	}

	return true
}

func bubbleSort(intervals []*Interval) []*Interval {

	for i := 0; i < len(intervals)-1; i++ {
		for j := i + 1; j < len(intervals); j++ {
			if intervals[i].Start > intervals[j].Start {
				intervals[i], intervals[j] = intervals[j], intervals[i]
			}
		}
	}

	return intervals
}