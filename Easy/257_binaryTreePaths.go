type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	var res []string

	if root == nil {
		return res
	}

	leaf(root, "", &res)
	return res
}

func leaf(node *TreeNode, str string, res *[]string) {
	str += strconv.Itoa(node.Val)
	if node.Left == nil && node.Right == nil {
		*res = append(*res, str)
		return
	}

	if node.Left != nil {
		leaf(node.Left, str+"->", res)
	}

	if node.Right != nil {
		leaf(node.Right, str+"->", res)
	}
}
