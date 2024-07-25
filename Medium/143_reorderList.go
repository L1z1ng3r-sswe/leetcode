/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	midNode := mid(head)

	half := reverse(midNode.Next)

	midNode.Next = nil

	for half != nil {
		tempFirst := head.Next
		tempSecond := half.Next

		head.Next = half
		half.Next = tempFirst

		head = tempFirst
		half = tempSecond
	}
}

func mid(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}

	return prev
}

// find mid
// reverse 2 part
// merge them

// time: O(N)
// space: O(1)