func maxPathSum(root *TreeNode) int {
	var maxPath = math.MinInt32

	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		leftRes := max(0, dfs(root.Left))
		rightRes := max(0, dfs(root.Right))

		total := root.Val + leftRes + rightRes

		maxPath = max(maxPath, total)

		return root.Val + max(leftRes, rightRes)
	}

	dfs(root)

	return maxPath
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// n - total number of nodes
// h - height of the bt
// time: O(n)
// space: O(h)