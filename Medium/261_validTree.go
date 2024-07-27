func ValidTree(n int, edges [][]int) bool {
	graph := make(map[int][]int)

	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	visited := make(map[int]bool)

	var dfs func(node int, parent int) bool
	dfs = func(node int, parent int) bool {
		visited[node] = true

		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				if !dfs(neighbor, node) {
					return false
				}
			} else if neighbor != parent {
				return false
			}
		}

		return true
	}

	if !dfs(0, -1) || len(visited) != n {
		return false
	}

	return true
}

// time: O(N+E)
// space: O(N+E)