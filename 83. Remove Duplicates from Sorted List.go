package main

func deleteDuplicates2(head *ListNode) *ListNode {
    if head == nil {
        return head
    }
    slow := head
    fast := head
    for fast != nil {
        for fast != nil && fast.Val == slow.Val {
            fast = fast.Next
        }
        slow.Next = fast
        slow = fast
    }
    return head
}
