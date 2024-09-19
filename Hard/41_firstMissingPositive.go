func firstMissingPositive(nums []int) int {
	var n = len(nums)

	for i := 0; i < n; i++ {
		for 0 < nums[i] && nums[i] < n && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}

	for i, num := range nums {
		if num != i+1 {
			return i + 1
		}
	}

	return n + 1
}

// time: O(n)
// space: O(1)