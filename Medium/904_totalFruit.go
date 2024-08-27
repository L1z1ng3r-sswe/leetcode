func totalFruit(fruits []int) int {
	var res int
	var left int
	window := make(map[int]int) // key - fruit, value - count

	for right := 0; right < len(fruits); right++ {
		rightFruit := fruits[right]
		window[rightFruit]++

		for len(window) > 2 {
			leftFruit := fruits[left]
			window[leftFruit]--
			if window[leftFruit] < 1 {
				delete(window, leftFruit)
			}
			left++
		}

		res = max(res, right-left+1)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// time: O(N)
// space: O(1)