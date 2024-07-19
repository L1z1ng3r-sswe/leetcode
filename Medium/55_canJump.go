func canJump(nums []int) bool {
	goal := len(nums) - 1

	for i := goal - 1; i >= 0; i-- {
		if i+nums[i] >= goal {
			goal = i
		}
	}

	if goal == 0 {
		return true
	}
	return false
}

// N = len(nums)
// time: O(N)
// space: O(1)