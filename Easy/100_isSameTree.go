


// bfs

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if q == nil || p == nil {
		return false
	}

	fifo1 := []*TreeNode{p}
	fifo2 := []*TreeNode{q}

	for len(fifo1) != 0 && len(fifo2) != 0 {
		node1 := fifo1[0]
		node2 := fifo2[0]
		fifo1 = fifo1[1:]
		fifo2 = fifo2[1:]

		if node1 == nil && node2 == nil {
			continue
		} else if node1 == nil || node2 == nil {
			return false
		} else if node1.Val != node2.Val {
			return false
		}

		fifo1 = append(fifo1, node1.Right, node1.Left)
		fifo2 = append(fifo2, node2.Right, node2.Left)
	}

	return true
}

// dfs

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}