func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if (targetSum-root.Val) == 0 && root.Right == nil && root.Left == nil {
		return true
	}

	return hasPathSum(root.Right, targetSum-root.Val) || hasPathSum(root.Left, targetSum-root.Val)
}