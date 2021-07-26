package main

func flatten(root *TreeNode) {
    if root == nil {
        return
    }
    dummyNode := &TreeNode{}
    cur := dummyNode
    stack := []*TreeNode{root}
    for len(stack) > 0 {
        root = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        cur.Right = root
        cur.Left = nil
        cur = cur.Right
        if root.Right != nil {
            stack = append(stack, root.Right)
        }
        if root.Left != nil {
            stack = append(stack, root.Left)
        }
    }
    root = dummyNode.Right
}
