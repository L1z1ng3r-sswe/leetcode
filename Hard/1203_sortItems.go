func sortItems(n int, m int, group []int, beforeItems [][]int) []int {
	groupNum := m

	for i, gr := range group {
		if gr == -1 {
			group[i] = groupNum
			groupNum++
		}
	}

	itemGraph := make([][]int, n)
	itemRel := make([]int, n)

	groupGraph := make([][]int, groupNum)
	groupRel := make([]int, groupNum)

	for node, roots := range beforeItems {
		for _, root := range roots {
			itemGraph[root] = append(itemGraph[root], node)
			itemRel[node]++

			if group[root] != group[node] {
				groupGraph[group[root]] = append(groupGraph[group[root]], group[node])
				groupRel[group[node]]++
			}
		}
	}

	sortedItems := topoSort(itemGraph, itemRel)
	if sortedItems == nil {
		return nil
	}
	sortedGroups := topoSort(groupGraph, groupRel)
	if sortedGroups == nil {
		return nil
	}

	groupMap := make(map[int][]int, groupNum)
	for _, item := range sortedItems {
		groupMap[group[item]] = append(groupMap[group[item]], item)
	}

	res := make([]int, 0, n)
	for _, gr := range sortedGroups {
		res = append(res, groupMap[gr]...)
	}

	return res
}

func topoSort(graph [][]int, rels []int) []int {
	queue := []int{}
	for node, rel := range rels {
		if rel == 0 {
			queue = append(queue, node)
		}
	}

	var res []int
	for len(queue) != 0 {
		root := queue[0]
		queue = queue[1:]
		res = append(res, root)

		for _, node := range graph[root] {
			rels[node]--
			if rels[node] == 0 {
				queue = append(queue, node)
			}
		}
	}

	if len(res) != len(graph) {
		return nil
	}

	return res
}

// n - number of nodes
// b - number of edges

// time: O(n+b)
// space: O(n+b)