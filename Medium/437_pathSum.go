func pathSum(root *TreeNode, targetSum int) int {
	var res int

	if root == nil {
		return res
	}

	var dfs func(root *TreeNode, currSum int)
	dfs = func(root *TreeNode, currSum int) {
		if currSum == targetSum {
			res++
		}

		if root.Left != nil {
			dfs(root.Left, currSum+root.Left.Val)
		}

		if root.Right != nil {
			dfs(root.Right, currSum+root.Right.Val)
		}
	}

	var traverse func(root *TreeNode)
	traverse = func(root *TreeNode) {
		dfs(root, root.Val)

		if root.Left != nil {
			traverse(root.Left)
		}

		if root.Right != nil {
			traverse(root.Right)
		}
	}

	traverse(root)

	return res
}

// n - total number of nodes
// time: O(n)
// space: stack