package main

func swapPairs(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    node := swapPairs(head.Next.Next)
    var pre *ListNode
    var next *ListNode
    cur := head
    for i := 0; i < 2; i++ {
        next = cur.Next
        cur.Next = pre
        pre = cur
        cur = next
    }
    head.Next = node
    return pre
}
