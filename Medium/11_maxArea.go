func maxArea(height []int) int {
	var res int
	left := 0
	right := len(height) - 1

	for left < right {
		res = max(res, (right-left)*min(height[left], height[right]))
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// time: O(n)
// space: O(1)