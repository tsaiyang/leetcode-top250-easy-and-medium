package main

func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil || k < 0 {
        return head
    }
    var tail *ListNode
    cur := head
    length := 0
    for cur != nil {
        if cur.Next == nil {
            tail = cur
        }
        length++
        cur = cur.Next
    }
    k = k % length
    if k == 0 {
        return head
    }
    var pre *ListNode
    cur = head
    for k >= 0 {
        pre = cur
        cur = cur.Next
        k--
    }
    pre.Next = nil
    tail.Next = head
    return cur
}
