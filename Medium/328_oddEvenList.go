func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	odd := head
	even := head.Next
	tmpEvenHead := even

	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}

	odd.Next = tmpEvenHead
	return head
}

// time: O(N)
// space: O(1)
