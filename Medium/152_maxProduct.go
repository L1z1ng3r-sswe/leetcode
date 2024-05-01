func maxProduct(nums []int) int {
	result := biggest(nums)
	maxRes, minRes := 1, 1

	for _, num := range nums {
		if num == 0 {
			maxRes, minRes = 1, 1
			continue
		}

		tmp := maxRes * num
		maxRes = biggest([]int{maxRes * num, minRes * num, num})
		minRes = smallest([]int{tmp, minRes * num, num})

		result = max(result, maxRes)
	}

	return result
}

func biggest(nums []int) int {
	result := nums[0]

	for i := 0; i < len(nums)-1; i++ {
		result = max(result, nums[i+1])
	}

	return result
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

func smallest(nums []int) int {
	result := nums[0]

	for i := 0; i < len(nums)-1; i++ {
		if result > nums[i+1] {
			result = nums[i+1]
		}
	}

	return result
}
