func removeNthFromEnd(head *ListNode, n int) *ListNode {
	reversed := reverse(head)

	dummy := &ListNode{Next: reversed}
	current := dummy
	counter := 0

	for current.Next != nil && counter < n {
		if counter == n-1 {
			current.Next = current.Next.Next
			break
		}
		current = current.Next
		counter++
	}

	return reverse(dummy.Next)
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}

