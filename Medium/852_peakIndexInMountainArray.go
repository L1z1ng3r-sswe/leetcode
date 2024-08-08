func peakIndexInMountainArray(arr []int) int {
	left := 0
	right := len(arr) - 1

	for left < right {
		mid := left + (right-left)/2

		if arr[mid] < arr[mid+1] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

// time: O(log(N))
// space: O(1)