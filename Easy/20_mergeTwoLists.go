func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	merged := &ListNode{0, nil}
	result := merged

	for list1 != nil && list2 != nil {

		if list1.Val > list2.Val {
			merged.Next = list2
			list2 = list2.Next
		} else {
			merged.Next = list1
			list1 = list1.Next
		}

		merged = merged.Next
	}

	for list1 != nil {
		merged.Next = list1
		list1 = list1.Next
		merged = merged.Next
	}

	for list2 != nil {
		merged.Next = list2
		list2 = list2.Next
		merged = merged.Next
	}

	return result.Next
}