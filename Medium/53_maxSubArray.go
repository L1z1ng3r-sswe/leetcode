func maxSubArray(nums []int) int {
	result := nums[0]
	current := 0

	for _, num := range nums {
		if current < 0 {
			current = 0
		}
		current += num
		result = max(result, current)
	}

	return result
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

