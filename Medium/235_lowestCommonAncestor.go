func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	pPath := targetPath(root, p)
	qPath := targetPath(root, q)

	var res *TreeNode

	for i, j := len(pPath)-1, len(qPath)-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if pPath[i] != qPath[j] {
			break
		}

		res = pPath[i]
	}

	return res
}

func targetPath(root *TreeNode, target *TreeNode) []*TreeNode {
	if root == nil {
		return nil
	}

	if root == target {
		return []*TreeNode{root}
	}

	if root.Left != nil {
		path := targetPath(root.Left, target)
		if path != nil {
			return append(path, root)
		}
	}

	if root.Right != nil {
		path := targetPath(root.Right, target)
		if path != nil {
			return append(path, root)
		}
	}
	return nil
}

// n - total number of nodes
// h - height of the tree
// time: O(n)
// space: O(h)