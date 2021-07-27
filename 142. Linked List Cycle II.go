package main

func detectCycle(head *ListNode) *ListNode {
    slow := head
    fast := head
    for fast != nil && fast.Next != nil {
        fast = fast.Next.Next
        slow = slow.Next
        if fast == slow {
            break
        }
    }
    fast = head
    for fast != slow {
        fast = fast.Next
        slow = slow.Next
    }
    return fast
}
