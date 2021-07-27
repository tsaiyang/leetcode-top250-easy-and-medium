package main

func preorderTraversal(root *TreeNode) []int {
    if root == nil {
        return nil
    }
    var result []int
    stack := []*TreeNode{root}
    for len(stack) > 0 {
        root = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        result = append(result, root.Val)
        if root.Right != nil {
            stack = append(stack, root.Right)
        }
        if root.Left != nil {
            stack = append(stack, root.Left)
        }
    }
    return result
}
