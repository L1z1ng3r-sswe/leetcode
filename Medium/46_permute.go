func permute(nums []int) [][]int {
	var result [][]int

	if len(nums) == 1 {
		return [][]int{append([]int{}, nums...)}
	}

	for i := 0; i < len(nums); i++ {
		n := nums[i]
		remaining := append([]int{}, nums[:i]...)
		remaining = append(remaining, nums[i+1:]...)

		perms := permute(remaining)

		for _, perm := range perms {
			perm = append(perm, n)
			result = append(result, perm)
		}
	}

	return result
}
