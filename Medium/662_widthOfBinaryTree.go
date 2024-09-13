type item struct {
	pos  int
	node *TreeNode
}

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var maxWidth int

	queue := []item{{0, root}}

	for len(queue) != 0 {
		queueLen := len(queue)
		maxWidth = max(maxWidth, queue[len(queue)-1].pos-queue[0].pos+1)

		for i := 0; i < queueLen; i++ {
			curr := queue[i]

			leftPos := curr.pos * 2
			rightPos := curr.pos*2 + 1

			if curr.node.Left != nil {
				queue = append(queue, item{leftPos, curr.node.Left})
			}

			if curr.node.Right != nil {
				queue = append(queue, item{rightPos, curr.node.Right})
			}
		}

		queue = queue[queueLen:]
	}

	return maxWidth
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// n - total number of nodes
// time: O(n)
// space: O(n)