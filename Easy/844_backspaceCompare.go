func backspaceCompare(s string, t string) bool {
	counters_index := len(s) - 1
	countert_index := len(t) - 1

	for i, j := len(s)-1, len(t)-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		s_index := currentIndex(counters_index, s)
		t_index := currentIndex(countert_index, t)

		if s_index < 0 && t_index < 0 {
			return true
		} else if s_index < 0 || t_index < 0 {
			return false
		} else if s[s_index] != t[t_index] {
			return false
		}

		counters_index = s_index - 1
		countert_index = t_index - 1

	}

	return true
}

func currentIndex(index int, s string) int {
	spaces := 0

	for index >= 0 {
		if s[index] == '#' {
			index--
			spaces++
		} else if spaces == 0 {
			break
		} else {
			spaces--
			index--
		}
	}

	return index
}
