func isValidBST(root *TreeNode) bool {
	var dfs func(root *TreeNode, min int, max int) bool
	dfs = func(root *TreeNode, min int, max int) bool {
		if root == nil {
			return true
		}

		//     left                right
		if root.Val >= min || root.Val <= max {
			return false
		}

		return dfs(root.Left, root.Val, max) && dfs(root.Right, min, root.Val)
	}

	return dfs(root, math.MaxInt64, math.MinInt64)
}

// n - total number of nodes
// time: O(n)
// space: stack