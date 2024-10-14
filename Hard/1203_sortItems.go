func sortItems(n int, m int, group []int, beforeItems [][]int) []int {
	groupId := m

	for i := 0; i < n; i++ {
		if group[i] == -1 {
			group[i] = groupId // if m = 2, then there are only 0 and 1 group, so -1 group becomes 2
			groupId++          // shift to the next group - 3
		}
	}

	itemGraph := make([][]int, n)
	itemRelations := make([]int, n)

	groupGraph := make([][]int, groupId)
	groupRelations := make([]int, groupId)

	// assemble graphs and relations
	for i := 0; i < n; i++ {
		for _, before := range beforeItems[i] {
			itemGraph[before] = append(itemGraph[before], i)
			itemRelations[i]++

			if group[before] != group[i] {
				groupGraph[group[before]] = append(groupGraph[group[before]], group[i])
				groupRelations[group[i]]++
			}
		}
	}

	groupOrder := topoSort(groupGraph, groupRelations)
	if groupOrder == nil {
		return []int{}
	}

	itemOrder := topoSort(itemGraph, itemRelations)
	if itemOrder == nil {
		return []int{}
	}

	groupToItems := make(map[int][]int)
	for _, item := range itemOrder {
		groupToItems[group[item]] = append(groupToItems[group[item]], item)
	}

	res := make([]int, 0, n)

	for _, grp := range groupOrder {
		res = append(res, groupToItems[grp]...)
	}

	return res
}

func topoSort(graph [][]int, relations []int) []int {
	n := len(graph)

	queue := []int{}

	for i := 0; i < n; i++ {
		if relations[i] == 0 {
			queue = append(queue, i)
		}
	}

	var res []int

	for len(queue) != 0 {
		root := queue[0]
		queue = queue[1:]
		res = append(res, root)

		for _, node := range graph[root] {
			relations[node]--
			if relations[node] == 0 { // this part checks on cycles
				queue = append(queue, node)
			}
		}
	}

	if len(res) != n {
		return nil
	}
	return res
}

// n - number of nodes
// b - number of edges

// time: O(n+b)
// space: O(n+b)