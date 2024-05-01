func maxArea(height []int) int {
	result := 0
	left := 0
	right := len(height) - 1

	for left != right {
		width := right - left
		area := width * min(height[left], height[right])
		result = max(result, area)

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return result
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

func min(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}