func partition(s string) [][]string {
	var res [][]string

	var dfs func(index int, substrings []string)
	dfs = func(index int, substrings []string) {
		if index == len(s) {
			res = append(res, append([]string(nil), substrings...))
			return
		}

		for i := index; i < len(s); i++ {
			if isPal(s, index, i) {
				dfs(i+1, append(substrings, s[index:i+1]))
			}
		}
	}
	dfs(0, []string{})

	return res
}

func isPal(str string, left, right int) bool {
	for left < right {
		if str[left] != str[right] {
			return false
		}
		left++
		right--
	}

	return true
}
