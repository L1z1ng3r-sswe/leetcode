func findMaxAverage(nums []int, k int) float64 {
	result := 0
	for i := 0; i < k; i++ {
		result += nums[i]
	}

	if len(nums)-k == 0 {
		return float64(result) / float64(k)
	}

	left := 1
	for i := len(nums) - k; i > 0; i-- {
		currSum := 0

		for right := left; right < left+k; right++ {
			num := nums[right]
			currSum += num
		}

		result = max(result, currSum)
		left++
	}

	return float64(result) / float64(k)
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}

	return n2
}
