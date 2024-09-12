type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	var res []int

	var descendants func(root *TreeNode, k int)
	descendants = func(root *TreeNode, k int) {
		if root == nil || k < 0 {
			return
		}
		if k == 0 {
			res = append(res, root.Val)
			return
		}
		descendants(root.Left, k-1)
		descendants(root.Right, k-1)
	}

	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return -1
		}

		if root == target {
			descendants(root, k)
			return 0
		}

		if root.Left != nil {
			leftDist := dfs(root.Left)
			if leftDist != -1 {
				descendants(root.Right, k-leftDist-2)
				return leftDist + 1
			}
		}

		if root.Right != nil {
			rightDist := dfs(root.Right)
			if rightDist != -1 {
				descendants(root.Left, k-rightDist-2)
				return rightDist + 1
			}
		}

		return -1
	}

	dfs(root)

	return res
}

// time: O(n)
// space: stack