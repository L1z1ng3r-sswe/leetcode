func maxProduct(nums []int) int {

	res := nums[0]
	high := nums[0]
	low := nums[0]

	for _, num := range nums[1:] {

		if num == 0 {
			high, low = 1, 1
			continue
		}

		tempHigh := high
		high = max(num, max(low*num, high*num))
		low = min(num, min(low*num, tempHigh))

		res = max(res, high)
	}

	return res
}
