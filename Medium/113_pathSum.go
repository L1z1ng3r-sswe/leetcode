func pathSum(root *TreeNode, targetSum int) [][]int {
	var res [][]int

	if root == nil {
		return res
	}

	var dfs func(root *TreeNode, currSum int, path []int)
	dfs = func(root *TreeNode, currSum int, path []int) {
		if root.Left == nil && root.Right == nil && currSum == targetSum {
			res = append(res, append([]int(nil), path...))
			return
		}

		if root.Left != nil {
			path = append(path, root.Left.Val)
			dfs(root.Left, currSum+root.Left.Val, path)
			path = path[:len(path)-1]
		}

		if root.Right != nil {
			path = append(path, root.Right.Val)
			dfs(root.Right, currSum+root.Right.Val, path)
			path = path[:len(path)-1]
		}
	}
	dfs(root, root.Val, []int{root.Val})

	return res
}