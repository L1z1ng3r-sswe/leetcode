package main

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// func GenerateLinkedList(arr []int) *ListNode {
// 	if len(arr) == 0 {
// 		return nil
// 	}

// 	head := &ListNode{Val: arr[0]}
// 	curr := head

// 	for i := 1; i < len(arr); i++ {
// 		curr.Next = &ListNode{Val: arr[i]}
// 		curr = curr.Next
// 	}

// 	return head
// }

// func PrintLinkedList(head *ListNode) {
// 	for head != nil {
// 		fmt.Printf("%d -> ", head.Val)
// 		head = head.Next
// 	}
// 	fmt.Println("nil")
// }

func main() {
	arr := []int{1, 2, 3, 4, 5}
	list := GenerateLinkedList(arr)

	node := reverseKGroup(list, 2)

	PrintLinkedList(node)
}

// infinity := []int{0}

// for i := 0; i < len(infinity); i++ {
// 	if i%60 == 0 {
// 		fmt.Printf("---%d---\n", i/60)
// 	} else {
// 		fmt.Println(i)
// 	}
// 	infinity = append(infinity, i+1)

// 	time.Sleep(time.Second)
// }

func reverseKGroup(head *ListNode, k int) *ListNode {
	node := head

	var count int
	for count < k && node != nil {
		node = node.Next
		count++
	}

	if count == k {
		reversedHead := reverse(head, k)
		head.Next = reverseKGroup(node, k)
		return reverseHead
	}

	return head
}

func reverse(head *ListNode, k int) *ListNode {
	var prev *ListNode
	curr := head

	for k > 0 {
		next := current.Next
		curr.Nesxt = prev
		prev = curr
		curr = next
		k--
	}

	return prev
}
