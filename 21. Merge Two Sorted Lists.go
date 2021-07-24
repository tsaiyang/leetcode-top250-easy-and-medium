package main

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    dummyNode := &ListNode{Val: 0}
    cur := dummyNode
    for l1 != nil || l2 != nil {
        if l1 == nil {
            cur.Next = l2
            cur = cur.Next
            l2 = l2.Next
            continue
        }
        if l2 == nil {
            cur.Next = l1
            cur = cur.Next
            l1 = l1.Next
            continue
        }
        if l1.Val < l2.Val {
            cur.Next = l1
            l1 = l1.Next
        } else {
            cur.Next = l2
            l2 = l2.Next
        }
        cur = cur.Next
    }
    return dummyNode.Next
}
