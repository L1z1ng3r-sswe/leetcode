func rob(nums []int) int {
	step1 := 0
	step2 := 0

	for _, num := range nums {
		temp := max(step1+num, step2)
		step1 = step2
		step2 = temp
	}

	return step2
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}