package main

type ListNode struct {
    Val  int
    Next *ListNode
}

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

type Node struct {
    Val   int
    Left  *Node
    Right *Node
    Next  *Node
}

type GraphNode struct {
    Val       int
    Neighbors []*GraphNode
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}
