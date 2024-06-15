func generateParenthesis(n int) []string {
	var res []string

	var dfs func(subStr string, open, cloth int)
	dfs = func(subStr string, open, cloth int) {
		if len(subStr) == n*2 {
			res = append(res, subStr)
			return
		}

		if open < n {
			dfs(subStr+"(", open+1, cloth)
		}
		if cloth < open {
			dfs(subStr+")", open, cloth+1)
		}
	}
	dfs("", 0, 0)

	return res
}