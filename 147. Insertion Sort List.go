package main

func insertionSortList(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    dummyHead := &ListNode{}
    dummyHead.Next = head
    head = head.Next
    cur := dummyHead.Next
    dummyHead.Next.Next = nil
    for head != nil {
        node := head
        var pre *ListNode
        cur = dummyHead.Next
        head = head.Next

        for cur != nil {
            if node.Val < cur.Val {
                break
            }
            pre = cur
            cur = cur.Next
        }

        if pre == nil {
            dummyHead.Next = node
            node.Next = cur
        } else {
            pre.Next = node
            node.Next = cur
        }
    }
    return dummyHead.Next
}
