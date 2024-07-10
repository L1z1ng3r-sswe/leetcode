func isSubtree(root, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}

	return isSameTree(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isSameTree(first, second *TreeNode) bool {
	if first == nil && second == nil {
		return true
	}

	if first == nil || second == nil {
		return false
	}

	return first.Val == second.Val && isSameTree(first.Left, second.Left) && isSameTree(first.Right, second.Right)
}
