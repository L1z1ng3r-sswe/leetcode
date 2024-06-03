func letterCasePermutation(s string) []string {
	var res []string

	var dfs func(currS string, index int)

	dfs = func(currS string, index int) {
		if len(currS) == len(s) {
			res = append(res, currS)
			return
		}

		var currChar = rune(s[index])

		if unicode.IsLetter(currChar) {
			dfs(currS+string(unicode.ToUpper(currChar)), index+1)
			dfs(currS+string(unicode.ToLower(currChar)), index+1)
		} else {
			dfs(currS+string(currChar), index+1)
		}
	}

	dfs("", 0)

	return res
}
