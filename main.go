package main

import "fmt"

func main() {
    fmt.Println(isIsomorphic("egg", "add"))
}

func printList(head *ListNode) {
    for head != nil {
        fmt.Println(head.Val)
        head = head.Next
    }
}

func generateList() *ListNode {
    l1 := &ListNode{Val: 1}
    l2 := &ListNode{Val: 2}
    l3 := &ListNode{Val: 3}
    l4 := &ListNode{Val: 4}
    l5 := &ListNode{Val: 5}
    l1.Next = l2
    l2.Next = l3
    l3.Next = l4
    l4.Next = l5
    return l1
}
