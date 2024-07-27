
func countComponents(n int, edges [][]int) int {
	graph := make(map[int][]int)
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	var visited = make(map[int]bool)

	var dfs func(node int)
	dfs = func(node int) {
		visited[node] = true

		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				dfs(neighbor)
			}
		}
	}

	var count int
	for i := 0; i < n; i++ {
		if !visited[i] {
			dfs(i)
			count++
		}
	}

	return count
}

// time: O(N+E)
// space: O(N+E)