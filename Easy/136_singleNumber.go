func singleNumber(nums []int) int {
	var mask int

	for _, num := range nums {
		mask ^= num
	}

	return mask
}

func singleNumber(nums []int) int {
	hashMap := make(map[int]bool)

	for _, num := range nums {
		if hashMap[num] {
			delete(hashMap, num)
		} else {
			hashMap[num] = true
		}
	}

	for key := range hashMap {
		return key
	}

	return 0
}
