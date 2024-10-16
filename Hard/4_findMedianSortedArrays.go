func findMedianSortedArrays(a []int, b []int) float64 {
	total := len(a) + len(b)
	half := total / 2

	if len(a) > len(b) {
		a, b = b, a
	}

	l, r := 0, len(a)

	for {
		i := (l + r) / 2 // midLeft
		j := half - i    // midRight

		aLeft := math.Inf(-1)
		if i > 0 {
			aLeft = float64(a[i-1])
		}

		aRight := math.Inf(1)
		if i < len(a) {
			aRight = float64(a[i])
		}

		bLeft := math.Inf(-1)
		if j > 0 {
			bLeft = float64(b[j-1])
		}

		bRight := math.Inf(1)
		if j < len(b) {
			bRight = float64(b[j])
		}

		if aLeft <= bRight && aRight >= bLeft {
			if total%2 == 0 {
				return (math.Max(aLeft, bLeft) + math.Min(aRight, bRight)) / 2 // even
			}
			return math.Min(aRight, bRight) // odd
		} else if aLeft > bRight {
			r = i - 1
		} else {
			l = i + 1
		}
	}
}

// time: O(log(m+n))