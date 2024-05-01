func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	size := 1
	fifo := []*TreeNode{root}

	for len(fifo) != 0 {
		lengthFifo := len(fifo)

		for i := 0; i < lengthFifo; i++ {
			tree := fifo[0]
			fifo = fifo[1:]

			if tree.Left == nil && tree.Right == nil {
				return size
			}

			if tree.Left != nil {
				fifo = append(fifo, tree.Left)
			}

			if tree.Right != nil {
				fifo = append(fifo, tree.Right)
			}
		}

		size++
	}
	return size
}