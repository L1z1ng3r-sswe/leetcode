func rob(nums []int) int {
	prevprev := 0
	prev := 0

	//[2,2000,9,3,1, 1000]
	for _, num := range nums {
		temp := max(prevprev+num, prev)
		prevprev = prev
		prev = temp
	}

	return prev
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}

	return n2
}