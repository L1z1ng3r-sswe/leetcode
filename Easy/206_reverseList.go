func reverseList(head *ListNode) *ListNode {
	var result *ListNode
	current := head

	for current != nil {
		next := current.Next
		current.Next = result
		result = current
		current = next
	}

	return result
}
