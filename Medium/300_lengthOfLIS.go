func lengthOfLIS(nums []int) int {
	dp := make(map[int]int)

	for i:=range nums {
		dp[i] = 1
	}

	for i:=range nums {
		for j:=0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j] + 1)
			}
		}
	}

	var res = 1
	for _, length := range dp {
		if res < length {
			res = length
		}
	}

	return res
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}

	return n2
}