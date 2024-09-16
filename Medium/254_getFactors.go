func getFactors(num int) [][]int {
	var factors [][]int

	if num <= 1 {
		return factors
	}

	var dfs func(currFactors []int, quotient, divisor int)
	dfs = func(currFactors []int, quotient, divisor int) {
		if quotient == 1 {
			factors = append(factors, append([]int(nil), currFactors...))
			return
		}

		for i := divisor; i <= quotient; i++ {
			if quotient%i == 0 {
				dfs(append(currFactors, i), quotient/i, i)
			}
		}
	}

	dfs([]int{}, num, 2)

	return factors
}

// time: O(n)
// space: stack