func letterCombinations(digits string) []string {
	var res []string

	if len(digits) == 0 {
		return res
	}

	var dfs func(index int, substr string)
	dfs = func(index int, substr string) {
		if len(substr) == len(digits) {
			res = append(res, substr)
			return
		}

		digit := string(digits[index])
		for _, char := range hs[digit] {
			dfs(index+1, substr+char)
		}
	}
	dfs(0, "")

	return res
}

var (
	hs = map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}
)