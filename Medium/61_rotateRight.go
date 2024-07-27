/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	var length int
	curr := head
	for curr != nil {
		curr = curr.Next
		length++
	}

	k = k % length
	if k == 0 {
		return head
	}

	newTail := head
	for i := 0; i < length-k-1; i++ {
		newTail = newTail.Next
	}
	newHead := newTail.Next

	newTail.Next = nil
	curr = newHead
	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = head

	return newHead
}

// time: O(N)
// space: O(1)