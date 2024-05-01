func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	prev1 := 1
	prev2 := 2

	for i := 3; i <= n; i++ {
		prev1, prev2 = prev2, prev1+prev2
	}

	return prev2
}
 
