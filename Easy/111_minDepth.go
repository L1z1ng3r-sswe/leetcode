func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	minDepth := 0
	fifo := []*TreeNode{root}

	for len(fifo) != 0 {
		length := len(fifo)
		minDepth++

		for i := 0; i != length; i++ {
			node := fifo[0]
			fifo = fifo[1:]

			if node.Right == nil && node.Left == nil {
				return minDepth
			}

			if node.Right != nil {
				fifo = append(fifo, node.Right)
			}

			if node.Left != nil {
				fifo = append(fifo, node.Left)
			}
		}
	}

	return minDepth
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	left := minDepth(root.Left)
	right := minDepth(root.Right)

	if root.Left == nil || root.Right == nil {
		return 1 + leftDepth + rightDepth
	}

	return min(left, right) + 1
}

func min(n1, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}