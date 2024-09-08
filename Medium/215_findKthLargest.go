func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(nums []int, left, right, k int) int {
	if left == right {
		return nums[left]
	}

	pivotIndex := left + rand.Intn(right-left+1)
	pivotIndex = partition(nums, left, right, pivotIndex)

	if pivotIndex == k {
		return nums[pivotIndex]
	} else if pivotIndex > k {
		return quickSelect(nums, left, pivotIndex-1, k)
	}
	return quickSelect(nums, pivotIndex+1, right, k)
}

func partition(nums []int, left, right, pivotIndex int) int {
	pivotValue := nums[pivotIndex]
	nums[right], nums[pivotIndex] = nums[pivotIndex], nums[right]
	currBigger := left

	for i := left; i < right; i++ {
		if nums[i] < pivotValue {
			nums[currBigger], nums[i] = nums[i], nums[currBigger]
			currBigger++
		}
	}

	nums[currBigger], nums[right] = nums[right], nums[currBigger]
	return currBigger
}

// space: O(~N)
// time: O(1)