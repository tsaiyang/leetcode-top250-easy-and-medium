package main

func levelOrder(root *TreeNode) [][]int {
    var result [][]int
    if root == nil {
        return result
    }
    queue := []*TreeNode{root}
    for len(queue) > 0 {
        size := len(queue)
        var arr []int
        for i := 0; i < size; i++ {
            node := queue[0]
            arr = append(arr, node.Val)
            queue = queue[1:]
            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
        result = append(result,arr)
    }
    return result
}
