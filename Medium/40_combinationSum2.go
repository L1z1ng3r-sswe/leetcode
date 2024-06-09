func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var res [][]int

	var dfs func(index int, currNums []int, remaining int)
	dfs = func(index int, currNums []int, remaining int) {
		if remaining == 0 {
			res = append(res, append([]int(nil), currNums...))
			return
		}
		if remaining < 0 {
			return
		}

		// [1,1,2,5,6,7,10]
		prev := -1
		for i := index; i < len(candidates); i++ {
			if prev == candidates[i] {
				continue
			}

			currNums = append(currNums, candidates[i])
			dfs(i+1, currNums, remaining-candidates[i])
			//pop
			currNums = currNums[:len(currNums)-1]
			prev = candidates[i]
		}
	}

	dfs(0, []int{}, target)

	return res
}