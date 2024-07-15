func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxProduct, minProduct := nums[0], nums[0]
	result := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			maxProduct, minProduct = minProduct, maxProduct
		}

		maxProduct = int(math.Max(float64(nums[i]), float64(maxProduct*nums[i])))
		minProduct = int(math.Min(float64(nums[i]), float64(minProduct*nums[i])))

		result = int(math.Max(float64(result), float64(maxProduct)))
	}

	return result
}
