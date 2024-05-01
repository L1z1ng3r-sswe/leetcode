func nextGreatestLetter(letters []byte, target byte) byte {
	var left, right, mid = 0, len(letters) - 1, 0

	for left <= right {
		mid = (left + right) / 2

		if letters[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if left < len(letters) {
		return letters[left]
	}
	return letters[0]
}     



  