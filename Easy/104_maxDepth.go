func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	result := 0
	fifo := []*TreeNode{root}

	for len(fifo) != 0 {
		length := len(fifo)
		result++

		for i := 0; i != length; i++ {
			node := fifo[0]
			fifo = fifo[1:]

			if node.Left != nil {
				fifo = append(fifo, node.Left)
			}

			if node.Right != nil {
				fifo = append(fifo, node.Right)
			}
		}
	}

	return result
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	return max(left, right) + 1
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}
