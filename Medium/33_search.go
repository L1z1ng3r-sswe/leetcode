func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] < nums[right] { // right is sorted
			if nums[mid] < target && target <= nums[right] { // target is in right
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else { // left is sorted
			if nums[left] <= target && target < nums[mid] { // target in left
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}

	return -1
}

// time: O(logN)
// space: O(1)