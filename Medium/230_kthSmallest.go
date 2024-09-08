func kthSmallest(root *TreeNode, k int) int {
	var res int
	var count = 0
	var dfs func(node *TreeNode)

	dfs = func(node *TreeNode) {
		if node == nil || count > k {
			return
		}

		dfs(node.Left)

		count++
		if count == k {
			res = node.Val
			return
		}

		dfs(node.Right)
	}
	dfs(root)

	return res
}

// time: O(logH)