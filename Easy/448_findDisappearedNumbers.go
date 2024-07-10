func findDisappearedNumbers(nums []int) []int {
	var disappeared []int

	i := 0
	for i < len(nums) {
		index := nums[i] - 1

		if nums[i] != nums[index] {
			nums[i], nums[index] = nums[index], nums[i]
		} else {
			i++
		}
	}

	for i, num := range nums {
		if num != i+1 {
			disappeared = append(disappeared, i+1)
		}
	}

	return disappeared
}

func findDisappearedNumbers(nums []int) []int {
	var disappeared []int

	for _, num := range nums {
		index := abs(num) - 1
		nums[index] = -abs(nums[index])
	}

	for i, num := range nums {
		if num > 0 {
			disappeared = append(disappeared, i+1)
		}
	}

	return disappeared
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}