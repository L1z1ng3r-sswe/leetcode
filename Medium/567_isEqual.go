func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	var s1Contain [26]int
	var s2Contain [26]int

	for i := 0; i < len(s1); i++ {
		s1Contain[s1[i]-'a']++
		s2Contain[s2[i]-'a']++
	}

	if isEqual(s1Contain, s2Contain) {
		return true
	}

	for right := len(s1); right < len(s2); right++ {
		left := right - len(s1)
		rightChar := s2[right] - 'a'
		leftChar := s2[left] - 'a'

		s2Contain[rightChar]++
		s2Contain[leftChar]--

		if isEqual(s1Contain, s2Contain) {
			return true
		}
	}

	return false
}

func isEqual(arr1, arr2 [26]int) bool {
	if len(arr1) != len(arr2) {
		return false
	}

	for i := 0; i < len(arr1); i++ {
		charA := arr1[i]
		charB := arr2[i]

		if charA != charB {
			return false
		}
	}

	return true
}

// time: O(N)
// space: O(1) // 'cause the array will have fixed length