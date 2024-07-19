func uniquePaths(m int, n int) int {
	mem := make([]int, n)
	for i := range mem {
		mem[i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			mem[j] += mem[j-1]
		}
	}

	return mem[n-1]
}

// time = m*n
// space = n