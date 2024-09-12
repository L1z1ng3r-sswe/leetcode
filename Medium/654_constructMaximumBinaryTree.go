func constructMaximumBinaryTree(nums []int) *TreeNode {
	var dfs func(start, end int) *TreeNode
	dfs = func(start, end int) *TreeNode {
		if start == end {
			return nil
		}

		pivot := findMaxIndex(nums, start, end)
		root := &TreeNode{Val: nums[pivot]}
		root.Left = dfs(start, pivot)
		root.Right = dfs(pivot+1, end)
		return root
	}

	return dfs(0, len(nums))
}

func findMaxIndex(nums []int, start, end int) int {
	pivot := start
	for i := start + 1; i < end; i++ {
		if nums[pivot] < nums[i] {
			pivot = i
		}
	}
	return pivot
}

// time: O(n^2)
// space: O(n)