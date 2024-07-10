func missingNumber(nums []int) int {
	n := len(nums)
	currentSum := n * (n + 1) / 2
	var initialSum int
	for _, num := range nums {
		initialSum += num
	}
	return currentSum - initialSum
}


