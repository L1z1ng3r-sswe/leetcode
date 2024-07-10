func twoSum(nums []int, target int) []int {
	hashTable := make(map[int]int)

	for i, num := range nums {
		if index, contains := hashTable[num]; contains {
			return []int{index, i}
		}
		hashTable[target-num] = i
	}

	return []int{}
}