func getFactors(num int) [][]int {
	var res [][]int

	var backtrack func(currFactors []int, num int, divisor int)
	backtrack = func(currFactors []int, num int, divisor int) {
		if num == 1 {
			res = append(res, append([]int(nil), currFactors))
			return
		}

		for i := divisor; i <= num; i++ {
			if num%i == 0 {
				backtrack(append(currFactors, i), num/i, i)
			}
		}
	}
	backtrack([]int{}, num, 2)

	return res
}