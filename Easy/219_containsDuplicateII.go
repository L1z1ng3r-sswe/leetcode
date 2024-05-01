func containsNearbyDuplicate(nums []int, k int) bool {
	hashMap := make(map[int]int)

	for index, num := range nums {
		if jndex, exist := hashMap[num]; exist && abs(index-jndex) <= k {
			return true
		} else {
			hashMap[num] = index
		}
	}

	return false
}

func abs(dif int) int {
	if dif < 0 {
		return -dif
	}
	return dif
}