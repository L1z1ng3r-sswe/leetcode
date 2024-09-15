func numSubarrayProductLessThanK(nums []int, k int) int {
	var res int
	var n = len(nums)
	var prod = 1

	left := 0

	for right := 0; right < n; right++ {
		prod *= nums[right]

		for prod >= k && left <= right {
			prod /= nums[left]
			left++
		}

		res += right - left + 1
	}

	return res
}

// time: O(n)
// space: O(1)