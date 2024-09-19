func reverseKGroup(head *ListNode, k int) *ListNode {
	node := head

	var count int
	for count < k && node != nil {
		node = node.Next
		count++
	}

	if count == k {
		reversed := reverse(head, k)
		head.Next = reverseKGroup(node, k)
		return reversed
	}

	return head
}

func reverse(head *ListNode, k int) *ListNode {
	var prev *ListNode
	curr := head

	for k > 0 {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
		k--
	}

	return prev
}

// time: O(n)
// space: stack