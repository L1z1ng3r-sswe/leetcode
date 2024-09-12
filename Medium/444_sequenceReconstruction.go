func sequenceReconstruction(org []int, seqs [][]int) bool {
	if len(org) == 0 && len(seqs) > 0 {
		return false
	}
	if len(org) > 0 && len(seqs) == 0 {
		return false
	}

	adjList := make(map[int][]int)
	inDegree := make(map[int]int)

	for _, num := range org {
		adjList[num] = []int{}
		inDegree[num] = 0
	}

	for _, seq := range seqs {
		for i := 0; i < len(seq); i++ {
			if _, exists := adjList[seq[i]]; !exists {
				return false
			}
			if i < len(seq)-1 {
				curr := seq[i]
				next := seq[i+1]
				adjList[curr] = append(adjList[curr], next)
				inDegree[next]++
			}
		}
	}

	queue := []int{}
	for node, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, node)
		}
	}

	var idx int
	for len(queue) > 0 {
		if len(queue) > 1 {
			return false
		}

		curr := queue[0]
		queue = queue[1:]

		if org[idx] != curr {
			return false
		}

		idx++

		for _, next := range adjList[curr] {
			inDegree[next]--
			if inDegree[next] == 0 {
				queue = append(queue, next)
			}
		}
	}

	return idx == len(org)
}

// n - total number of nums in sequences
// time: O(n)
// space: O(n)