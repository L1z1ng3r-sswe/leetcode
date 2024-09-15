func threeSumClosest(nums []int, target int) int {
	var res = math.MaxInt64
	var n = len(nums)

	sort.Ints(nums)

	for i, num := range nums[:n-1] {
		left := i + 1
		right := n - 1

		for left < right {
			currSum := num + nums[left] + nums[right]
			if abs(target-currSum) < abs(target-res) {
				res = currSum
			}

			if currSum < target {
				left++
			} else if target < currSum {
				right--
			} else {
				return currSum
			}
		}
	}

	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// time: O(n^2)
// space: O(1)