func maxSubArray(nums []int) int {
	var res = nums[0]
	var currentMax = nums[0]

	for _, num := range nums {
		currentMax = max(currentMax+num, num)
		res = max(res, currentMax)
	}

	return res
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

