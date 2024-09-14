func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	rootIdx := findElem(preorder[0], inorder)

	node := &TreeNode{
		Val: preorder[0],
	}
	node.Left = buildTree(preorder[1:1+rootIdx], inorder[:rootIdx])
	node.Right = buildTree(preorder[rootIdx+1:], inorder[rootIdx+1:])

	return node
}

func findElem(target int, nums []int) int {
	for i, num := range nums {
		if num == target {
			return i
		}
	}

	return -1
}

// time: O(n^2)
// space: stack