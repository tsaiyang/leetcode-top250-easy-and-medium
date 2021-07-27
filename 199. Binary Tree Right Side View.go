package main

func rightSideView(root *TreeNode) []int {
    if root == nil {
        return nil
    }
    var result []int
    queue := []*TreeNode{root}
    for len(queue) > 0 {
        size := len(queue)
        for i := 0; i < size; i++ {
            root = queue[0]
            queue = queue[1:]
            if i == size-1 {
                result = append(result, root.Val)
            }
            if root.Left != nil {
                queue = append(queue, root.Left)
            }
            if root.Right != nil {
                queue = append(queue, root.Right)
            }
        }
    }
    return result
}
