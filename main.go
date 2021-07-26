package main

import "fmt"

func main() {
    fmt.Println(restoreIpAddresses("1111"))
}

func generateList() *ListNode {
    l1 := &ListNode{Val: 1}
    l2 := &ListNode{Val: 2}
    l3 := &ListNode{Val: 3}
    l4 := &ListNode{Val: 3}
    l5 := &ListNode{Val: 4}
    l6 := &ListNode{Val: 4}
    l7 := &ListNode{Val: 5}

    l1.Next = l2
    l2.Next = l3
    l3.Next = l4
    l4.Next = l5
    l5.Next = l6
    l6.Next = l7
    return l1
}

func printList(head *ListNode) {
    for head != nil {
        fmt.Println(head.Val)
        head = head.Next
    }
}
