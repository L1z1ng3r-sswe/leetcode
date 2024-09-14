type nodePos struct {
	node *TreeNode
	pos  int
}

func widthOfBinaryTree(root *TreeNode) int {
	queue := []nodePos{{node: root, pos: 0}}
	maxWidth := 0

	for len(queue) > 0 {
		levelSize := len(queue)
		firstPos := queue[0].pos
		lastPos := queue[len(queue)-1].pos
		maxWidth = max(maxWidth, lastPos-firstPos+1)

		for i := 0; i < levelSize; i++ {
			currNode := queue[0].node
			currPos := queue[0].pos
			queue = queue[1:]

			if currNode.Left != nil {
				queue = append(queue, nodePos{node: currNode.Left, pos: 2 * currPos})
			}
			if currNode.Right != nil {
				queue = append(queue, nodePos{node: currNode.Right, pos: 2*currPos + 1})
			}
		}
	}
	return maxWidth
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// time: O(n)
// space: O(n)