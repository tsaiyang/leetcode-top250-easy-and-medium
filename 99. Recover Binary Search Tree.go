package main

func recoverTree(root *TreeNode) {
    if root == nil {
        return
    }
    var node1, node2, pre *TreeNode
    var stack []*TreeNode
    for len(stack) > 0 || root != nil {
        if root != nil {
            stack = append(stack, root)
            root = root.Left
        } else {
            root = stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            if pre != nil && root.Val < pre.Val {
                if node1 == nil {
                    node1 = pre
                }
                node2 = root
            }
            pre = root
            root = root.Right
        }
    }
    node1.Val, node2.Val = node2.Val, node1.Val
}
