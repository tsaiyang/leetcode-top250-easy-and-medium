package main

func removeElements(head *ListNode, val int) *ListNode {
    if head == nil {
        return nil
    }
    dummyHead := &ListNode{}
    dummyHead.Next = head
    cur := dummyHead
    for head != nil {
        if head.Val == val {
            cur.Next = head.Next
            head = cur.Next
        } else {
            cur = head
            head = head.Next
        }
    }
    return dummyHead.Next
}
