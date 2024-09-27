func minKBitFlips(nums []int, k int) int {
	var n = len(nums)
	flipStore := make([]int, n)
	var flip int
	var res int

	for i := 0; i < n; i++ {
		if i >= k {
			flip ^= flipStore[i-k]
		}

		if flip^nums[i] == 0 {
			if i+k > n {
				return -1
			}

			flip ^= 1
			flipStore[i] = 1
			res++
		}
	}

	return res
}

// time: O(n)
// space: O(n)