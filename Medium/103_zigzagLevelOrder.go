func zigzagLevelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	var queue = []*TreeNode{root}
	var isToLeft bool

	for len(queue) != 0 {
		queueLen := len(queue)
		level := []int{}

		for i := 0; i < queueLen; i++ {
			node := queue[i]

			if !isToLeft {
				level = append(level, node.Val)
			} else {
				level = append(level, queue[queueLen-i-1].Val)
			}

			if !isToLeft {
			} else {

			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, level)
		isToLeft = !isToLeft

		queue = queue[queueLen:]
	}

	return res
}

// time: O(n)
// space: O(n)