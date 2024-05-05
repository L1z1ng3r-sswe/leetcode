
 
	type TreeNode struct {
		Val   int
		Left  *TreeNode
		Right *TreeNode
	}

	func binaryTreePaths(root *TreeNode) []string {
		if root == nil {
			return []string{}
		}

		var result []string
		delve(root, "", &result)
		return result
	}

	func delve(node *TreeNode, previous string, result *[]string) {
		previous += strconv.Itoa(node.Val)
		if node.Right == nil && node.Left == nil {
			*result = append(*result, previous)
			return
		}

		if node.Left != nil {
			delve(node.Left, previous+"->", result)
		}

		if node.Right != nil {
			delve(node.Right, previous+"->", result)
		}
	}