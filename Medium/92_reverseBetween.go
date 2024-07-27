func reverseBetween(head *ListNode, left int, right int) *ListNode {
  var dummy = &ListNode{Next: head} // dummy for edge cases 
  currPrev := dummy // value before left

  for i:=0; i < left-1; i++ { 
    currPrev = currPrev.Next 
  }
  
  curr := currPrev.Next // left

  var prev *ListNode // nil
  // 1 -> 2 -> 3 -> 4 -> 5 -> nil
  for i:=0; i < right-left+1; i++ {
    temp := curr.Next // 2 -> 3 -> 4 -> 5 -> nil
    curr.Next = prev // 1 -> nil
    prev = curr // 1 -> nil
    curr = temp // 2 -> 3 -> 4 -> 5 -> nil
  }

  currPrev.Next.Next = curr
  currPrev.Next = prev

  return dummy.Next
}



// time:
// space:

git add .

git commit -m "feat: 92_countComponents done; time: O(), space: O()"

git push origin main

