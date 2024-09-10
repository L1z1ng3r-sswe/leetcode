func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	allNodes := make([][]int, n)
	degree := make([]int, n)

	for _, edge := range edges {
		root := edge[0]
		leaf := edge[1]
		allNodes[root] = append(allNodes[root], leaf)
		allNodes[leaf] = append(allNodes[leaf], root)
		degree[root]++
		degree[leaf]++
	}

	leaves := []int{}
	for i := 0; i < n; i++ {
		if len(allNodes[i]) == 1 {
			leaves = append(leaves, i)
		}
	}

	nodes := n
	for nodes > 2 {
		nodes -= len(leaves)
		newLeaves := []int{}

		for _, leaf := range leaves {
			for _, root := range allNodes[leaf] {
				degree[root]--

				if degree[root] == 1 {
					newLeaves = append(newLeaves, root)
				}
			}
		}
		leaves = newLeaves
	}

	return leaves
}

// time: O(N)
// space: O(N)