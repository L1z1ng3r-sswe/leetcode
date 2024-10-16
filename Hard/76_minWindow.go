func minWindow(s string, t string) string {
	var res string

	tFreq := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		tFreq[t[i]]++
	}

	window := make(map[byte]int)
	count := len(tFreq)
	left := 0

	for right := 0; right < len(s); right++ {
		ch := s[right]

		if _, ok := tFreq[ch]; ok {
			window[ch]++

			if window[ch] == tFreq[ch] {
				count--

				for count == 0 {
					if len(res) == 0 || len(res) > len(s[left:right+1]) {
						res = s[left : right+1]
					}

					leftChar := s[left]
					left++

					if _, ok := tFreq[leftChar]; ok {
						window[leftChar]--
						if window[leftChar] < tFreq[leftChar] {
							count++
						}
					}
				}
			}
		}
	}

	return res
}

// time: O(n)
// space: O(n)