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

	shift := k % length

	if shift == 0 {
		return head
	}

	tail := head

	// tail = 1 -> 2 -> 3 -> 4 -> 5 -> nil
	for i := 0; i < length-shift-1; i++ {
		tail = tail.Next
	}

	// newHead = 4 -> 5 -> nil
	newHead := tail.Next

	// head = 1 -> 2 -> 3 -> nil
	tail.Next = nil

	// tail = 3 -> 4 -> 5 -> nil
	result := newHead
	for result.Next != nil {
		result = result.Next
	}

	result.Next = head

	return newHead
}

// time: O(N)
// space: O(1)

