func minSubArrayLen(target int, nums []int) int {
	var res = math.MaxInt64
	curr := 0
	left := 0

	for right := 0; right < len(nums); right++ {
		curr += nums[right]

		for curr >= target {
			res = min(res, right-left+1)
			curr -= nums[left]
			left++
		}
	}

	if res == math.MaxInt64 {
		return 0
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// time: O(N^2)
// space: O(1)