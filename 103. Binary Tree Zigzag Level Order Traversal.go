package main

func zigzagLevelOrder(root *TreeNode) [][]int {
    var result [][]int
    if root == nil {
        return result
    }
    deque := []*TreeNode{root}
    flag := true
    for len(deque) > 0 {
        size := len(deque)
        var arr []int
        for i := 0; i < size; i++ {
            if flag {
                node := deque[0]
                deque = deque[1:]
                arr = append(arr, node.Val)
                if node.Left != nil {
                    deque = append(deque, node.Left)
                }
                if node.Right != nil {
                    deque = append(deque, node.Right)
                }
            } else {
                node := deque[len(deque)-1]
                deque = deque[:len(deque)-1]
                arr = append(arr, node.Val)
                if node.Right != nil {
                    deque = append([]*TreeNode{node.Right}, deque...)
                }
                if node.Left != nil {
                    deque = append([]*TreeNode{node.Left}, deque...)
                }
            }
        }
        flag = !flag
        result = append(result, arr)
    }
    return result
}
