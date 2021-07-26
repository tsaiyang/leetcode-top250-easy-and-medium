package main

func reverseBetween(head *ListNode, left int, right int) *ListNode {
    var pre *ListNode
    var post *ListNode
    cur := head
    count := 0
    for cur != nil {
        count++
        if left-1 == count {
            pre = cur
        }
        if right+1 == count {
            post = cur
        }
        cur = cur.Next
        count++
    }
    var n1 *ListNode
    if pre == nil {
        n1 = head
    } else {
        n1 = pre.Next
    }
    n1.Next = post
    n2 := n1.Next
    var next *ListNode
    for n2 != post {
        next = n2.Next
        n2.Next = n1
        n1 = n2
        n2 = next
    }
    if pre == nil {
        head = n1
    }else{
        pre.Next = n1
    }
    return head
}
