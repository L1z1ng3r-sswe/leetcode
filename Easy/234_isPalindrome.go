func isPalindrome(head *ListNode) bool {
	middle := mid(head)
	reversed := reverse(middle)

	for reversed != nil {
		if head.Val != reversed.Val {
			return false
		}
		head = head.Next
		reversed = reversed.Next
	}

	return true
}

func mid(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}

func reverse(head *ListNode) *ListNode {
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