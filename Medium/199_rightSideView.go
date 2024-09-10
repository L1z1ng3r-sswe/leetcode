/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	var queue = []*TreeNode{root}

	for len(queue) != 0 {
		res = append(res, queue[len(queue)-1].Val)
		queueLen := len(queue)

		for i := 0; i < queueLen; i++ {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return res
}

// time: O(n)
// space: O(w)