func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	mid := getMid(head)
	left := head
	right := mid.Next
	mid.Next = nil

	left = sortList(left)
	right = sortList(right)

	return merge(left, right)
}

// 1 > 2 > 3 > nil
func getMid(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func merge(nums1 *ListNode, nums2 *ListNode) *ListNode {
	dummy := &ListNode{}
	sorted := dummy

	for nums1 != nil && nums2 != nil {
		if nums1.Val < nums2.Val {
			sorted.Next = nums1
			nums1 = nums1.Next
		} else {
			sorted.Next = nums2
			nums2 = nums2.Next
		}
		sorted = sorted.Next
	}

	if nums1 != nil {
		sorted.Next = nums1
	} else if nums2 != nil {
		sorted.Next = nums2
	}

	return dummy.Next
}