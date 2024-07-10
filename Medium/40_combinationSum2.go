func combinationSum2(candidates []int, target int) [][]int {
	var res [][]int
	sort.Ints(candidates)

	var dfs func(start int, currNums []int, remaining int)
	dfs = func(start int, currNums []int, remaining int) {
		if remaining == 0 {
			res = append(res, append([]int(nil), currNums...))
			return
		}

		if remaining < 0 {
			return
		}

		// [1,1,1,2,2], target = 4
		var prev = -1
		for i := start; i < len(candidates); i++ {
			if prev == candidates[i] {
				continue
			}

			currNums = append(currNums, candidates[i])
			dfs(i+1, currNums, remaining-candidates[i])
			currNums = currNums[:len(currNums)-1]
			prev = candidates[i]
		}
	}
	dfs(0, []int{}, target)

	return res
}
