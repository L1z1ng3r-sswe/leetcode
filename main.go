

func permuteUnique(nums []int) [][]int {
	var res [][]int
	var ht = make(map[int]bool)

	for _, num := range nums {
		ht[num] = false
	}

	if len(nums) == 1 {
		return [][]int{nums}
	}

	for i := 0; i < len(nums); i++ {
		if ht[nums[i]] {
			continue
		}

		ht[nums[i]] = true

		pop := nums[i]
		poped := append([]int(nil), nums[:i]...)
		poped = append(poped, nums[i+1:]...)

		subRes := permuteUnique(poped)
		for _, poped := range subRes {
			poped = append(poped, pop)
			res = append(res, poped)
		}
	}

	return res
}
