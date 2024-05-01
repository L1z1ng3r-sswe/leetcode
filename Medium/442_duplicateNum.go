func duplicateNum(nums []int) []int {
	var duplicate []int

	for _, num := range nums {
		index := abs(num) - 1

		if nums[index] < 0 {
			duplicate = append(duplicate, abs(num))
		} else {
			nums[index] = -nums[index]
		}
	}

	return duplicate
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

