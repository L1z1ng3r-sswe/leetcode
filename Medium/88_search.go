func search(nums []int, target int) bool {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return true
		}

		if nums[left] == nums[mid] && nums[mid] == nums[right] {
			left++
			right--
		} else if nums[mid] <= nums[right] { // right is sorted
			if nums[mid] < target && target <= nums[right] { // if target in the right
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else { // left is sorted
			if nums[left] <= target && target < nums[mid] { // if target in the left
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}

	return false
}

// time := O(logN)
// space := O(1)

