func productExceptSelf(nums []int) []int {
	firstArr := make([]int, len(nums))
	secondArr := make([]int, len(nums))

	before := 1

	for i := 0; i < len(nums); i++ {
		firstArr[i] = before
		before = before * nums[i]
	}

	after := 1
	for i := len(nums) - 1; i >= 0; i-- {
		secondArr[i] = firstArr[i] * after
		after *= nums[i]
	}

	return secondArr
}