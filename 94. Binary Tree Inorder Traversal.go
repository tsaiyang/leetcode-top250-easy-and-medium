package main

func inorderTraversal(root *TreeNode) []int {
    var stack []*TreeNode
    var result []int
    for len(stack) > 0 || root != nil {
        if root != nil {
            stack = append(stack, root)
            root = root.Left
        } else {
            root = stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            result = append(result, root.Val)
            root = root.Right
        }
    }
    return result
}
