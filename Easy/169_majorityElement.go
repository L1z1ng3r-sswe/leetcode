func majorityElement(nums []int) int {
	result := 0
	vote := 0

	for _, num := range nums {
		if vote == 0 {
			result = num
		} else if result != num {
			vote--
			continue
		}

		vote++
	}

	return result
}

