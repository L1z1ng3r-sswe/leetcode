func findDuplicate(nums []int) int {
	turtle := 0
	rabbit := 0

	for {
		turtle = nums[turtle]
		rabbit = nums[nums[rabbit]]
		if turtle == rabbit {
			break
		}
	}

	turtle = 0
	for turtle != rabbit {
		rabbit = nums[rabbit]
		turtle = nums[turtle]
	}

	return turtle
}

// 2 <= len(nums)
// 1 <= nums[i] <= n
// total repeat -> eternal