package main

func partition(head *ListNode, x int) *ListNode {
    lessHead := &ListNode{}
    largerHead := &ListNode{}
    lessCur := lessHead
    largerCur := largerHead
    cur := head
    for cur != nil {
        if cur.Val < x {
            lessCur.Next = cur
            lessCur = lessCur.Next
        } else {
            largerCur.Next = cur
            largerCur = largerCur.Next
        }
        cur = cur.Next
    }
    if largerCur != nil {
        largerCur.Next = nil
    }
    if lessHead.Next == nil {
        return largerHead.Next
    }
    lessCur.Next = largerHead.Next
    return lessHead.Next
}
