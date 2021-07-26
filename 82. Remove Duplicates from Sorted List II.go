package main

func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return head
    }
    dummyNode := &ListNode{}
    dummyNode.Next = head
    pre := dummyNode
    slow := head
    fast := head
    count := 0
    for fast != nil {
        for fast != nil && fast.Val == slow.Val {
            fast = fast.Next
            count++
        }
        if count > 1 {
            pre.Next = fast
            slow = fast
        } else {
            pre = slow
            slow = fast
        }
        count = 0
    }
    return dummyNode.Next
}
