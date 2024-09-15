func threeSum(nums []int) [][]int {
	var n = len(nums)
	var res [][]int

	sort.Ints(nums)

	for i, num := range nums[:n-1] {
		if i > 0 && num == nums[i-1] {
			continue
		}
		left := i + 1
		right := n - 1

		for left < right {
			currSum := num + nums[left] + nums[right]

			if currSum < 0 {
				left++
			} else if currSum > 0 {
				right--
			} else {
				res = append(res, []int{num, nums[left], nums[right]})
				left++
				right--

				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			}
		}
	}

	return res
}

// time: O(n^2)
// space: O(1)