

// BFS

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	fifo := []*TreeNode{root}
	sum := []int{targetSum - root.Val}

	for len(fifo) != 0 {
		node := fifo[0]
		currentSum := sum[0]

		fifo = fifo[1:]
		sum = sum[1:]

		if currentSum == 0 && node.Left == nil && node.Right == nil {
			return true
		}

		if node.Left != nil {
			fifo = append(fifo, node.Left)
			sum = append(sum, currentSum-node.Left.Val)
		}

		if node.Right != nil {
			fifo = append(fifo, node.Right)
			sum = append(sum, currentSum-node.Right.Val)
		}
	}

	return false
}

// DFS

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	if targetSum == root.Val && root.Left == nil && root.Right == nil {
		return true
	}

	return hasPathSum(root.Right, targetSum-root.Val) || hasPathSum(root.Left, targetSum-root.Val)
}