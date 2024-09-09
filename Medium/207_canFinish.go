func canFinish(numCourses int, prerequisites [][]int) bool {
	coursesAllowedAfterCompletingOne := make([][]int, numCourses)
	courseUsage := make([]int, numCourses)

	for _, prerequisite := range prerequisites {
		before := prerequisite[1]
		after := prerequisite[0]

		coursesAllowedAfterCompletingOne[before] = append(coursesAllowedAfterCompletingOne[before], after)
		courseUsage[after]++
	}

	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if courseUsage[i] == 0 {
			queue = append(queue, i)
		}
	}

	var count int
	for len(queue) != 0 {
		curr := queue[0]
		queue = queue[1:]
		count++

		for _, toComplete := range coursesAllowedAfterCompletingOne[curr] {
			courseUsage[toComplete]--
			if courseUsage[toComplete] == 0 {
				queue = append(queue, toComplete)
			}
		}
	}

	return count == numCourses
}

// time: O(V+E) V - the number of courses, E - the number of prerequisites