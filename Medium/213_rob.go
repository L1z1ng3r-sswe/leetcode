func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	prev2 := 0
	prev1 := 0

	for i := len(nums) - 1; i > 0; i-- {
		curr := nums[i]
		maxSoFar := max(prev2+curr, prev1)
		prev2 = prev1
		prev1 = maxSoFar
	}

	maxClockwise := robClockwise(nums)
	return max(prev1, maxClockwise)
}

func robClockwise(nums []int) int {
	prev2 := 0
	prev1 := 0

	for i := 0; i < len(nums)-1; i++ {
		curr := nums[i]
		maxSoFar := max(prev2+curr, prev1)
		prev2 = prev1
		prev1 = maxSoFar
	}

	return prev1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
