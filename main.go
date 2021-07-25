package main

import "fmt"

func main() {
    //b1 := []byte{'5', '3', ' ', ' ', '7', ' ', ' ', ' ', ' '}
    //b2 := []byte{'6', ' ', ' ', '1', '9', '5', ' ', ' ', ' '}
    //b3 := []byte{' ', '9', '8', ' ', ' ', ' ', ' ', '6', ' '}
    //b4 := []byte{'8', ' ', ' ', ' ', '6', ' ', ' ', ' ', '3'}
    //b5 := []byte{'4', ' ', ' ', '8', ' ', '3', ' ', ' ', '1'}
    //b6 := []byte{'7', ' ', ' ', ' ', '2', ' ', ' ', ' ', '6'}
    //b7 := []byte{' ', '6', ' ', ' ', ' ', ' ', '2', '8', ' '}
    //b8 := []byte{' ', ' ', ' ', '4', '1', '9', ' ', ' ', '5'}
    //b9 := []byte{' ', ' ', ' ', ' ', '8', ' ', ' ', '7', '9'}
    //board := [][]byte{b1, b2, b3, b4, b5, b6, b7, b8, b9}
    //fmt.Println(isValidSudoku(board))
    //fmt.Println(generateMatrix(4))
    head := generateList()
    head = rotateRight(head, 1)
    printList(head)
}

func generateList() *ListNode {
    l1 := &ListNode{Val: 1}
    l2 := &ListNode{Val: 2}
    l1.Next = l2
    return l1
}

func printList(head *ListNode) {
    for head != nil {
        fmt.Println(head.Val)
        head = head.Next
    }
}
