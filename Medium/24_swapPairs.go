func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{
		Next: head,
	}
	prev := dummy

	for head != nil && head.Next != nil {
		first := head
		second := head.Next

		prev.Next = second
		first.Next = second.Next
		second.Next = first

		prev = first
		head = first.Next
	}

	return dummy.Next
}

// time: O(N)
// space: O(1)