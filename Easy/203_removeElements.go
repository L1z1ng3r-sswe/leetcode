func removeElements(head *ListNode, val int) *ListNode {
	result := &ListNode{0, head}
	current := result

	for current.Next != nil {
		if current.Next.Val == val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	return result.Next
}