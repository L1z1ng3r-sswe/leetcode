func findNumberOfLIS(nums []int) int {
	var maxLen = 1
	var maxCount int
	var lens = make([]int, len(nums))
	var counts = make([]int, len(nums))

	for i := range nums {
		lens[i] = 1
		counts[i] = 1
	}

	for i := range nums {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if lens[j]+1 > lens[i] {
					lens[i] = lens[j] + 1
					counts[i] = counts[j]
				} else if lens[j]+1 == lens[i] {
					counts[i] += counts[j]
				}
			}
		}

		maxLen = max(maxLen, lens[i])
	}

	for i, len := range lens {
		if len == maxLen {
			maxCount += counts[i]
		}
	}

	return maxCount
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

// N = len(nums)
// time: O(N**N)
// space: O(N)