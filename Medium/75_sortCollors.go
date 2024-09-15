func sortColors(nums []int) {
	low := 0
	mid := 0
	high := len(nums) - 1

	for mid <= high {
		switch nums[mid] {
		case 0:
			nums[low], nums[mid] = nums[mid], nums[low]
			low++
			mid++
		case 1:
			mid++
		case 2:
			nums[high], nums[mid] = nums[mid], nums[high]
			high--
		}
	}
}

// time: O(n)
// space: O(1)