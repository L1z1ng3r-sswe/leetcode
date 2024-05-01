func minSubArrayLen(target int, nums []int) int {
	result, total := math.MaxInt64, 0
	left := 0

	for right, num := range nums {
		total += num

		for total >= target {
			result = min(result, right-left+1)
			total -= nums[left]
			left++
		}
	}

	if result == math.MaxInt64 {
		return 0
	}
	return result
}

func min(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}