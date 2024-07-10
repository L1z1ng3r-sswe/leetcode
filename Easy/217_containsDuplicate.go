func containsDuplicate(nums []int) bool {
	hashMap := make(map[int]bool)

	for _, num := range nums {
		if hashElem := hashMap[num]; hashElem {
			return true
		} else {
			hashMap[num] = true
		}
	}

	return false
}