func isBalanced(root *TreeNode) bool {
	_, res := dfs(root)
	return res
}

func dfs(node *TreeNode) (int, bool) {
	if node == nil {
		return 0, true
	}

	balanceLeft, isBalancedLeft := dfs(node.Left)
	balanceRight, isBalancedRight := dfs(node.Right)

	if !isBalancedLeft || !isBalancedRight {
		return 0, false
	}

	if abs(balanceLeft-balanceRight) > 1 {
		return 0, false
	}

	return 1 + max(balanceLeft, balanceRight), true
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// time: O(N)
// space: O(N) - due to recursion call stack, but we can round it to O(1)