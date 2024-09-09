func findOrder(numCourses int, prerequisites [][]int) []int {
	adjList := make([][]int, numCourses)
	availableCourses := make([]int, numCourses)

	for _, prerequisite := range prerequisites {
		before := prerequisite[1]
		after := prerequisite[0]
		adjList[before] = append(adjList[before], after)
		availableCourses[after]++
	}

	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if availableCourses[i] == 0 {
			queue = append(queue, i)
		}
	}

	var res []int
	for len(queue) != 0 {
		curr := queue[0]
		queue = queue[1:]
		res = append(res, curr)

		for _, availableCourse := range adjList[curr] {
			availableCourses[availableCourse]--
			if availableCourses[availableCourse] == 0 {
				queue = append(queue, availableCourse)
			}
		}
	}

	if len(res) != numCourses {
		return []int{}
	}

	return res
}

// V - number of courses, E - number of prerequisites
// time: O(V + E)
// space: O(V + E)