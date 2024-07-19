func countSubstrings(s string) int {
	var res int

	var dfs func(left, right int)
	dfs = func(left, right int) {
		for left >= 0 && right < len(s) && s[left] == s[right] {
			left--
			right++

			res++
		}
	}

	for i := 0; i < len(s); i++ {
		//for cases like: aaa
		dfs(i, i)

		//for cases like: aa
		dfs(i, i+1)
	}

	return res
}

// N = len(s)
// time: O(N**N)