func minSubArrayLen(target int, nums []int) int {
	var res = math.MaxInt
	var left int

	for right := 0; right < len(nums); right++ {
		rightNum := nums[right]
		target -= rightNum

		for target <= 0 {
			res = min(res, right-left+1)
			leftNum := nums[left]
			target += leftNum
			left++
		}
	}

	if res == math.MaxInt {
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

// time: O(N)
// space: O(1)