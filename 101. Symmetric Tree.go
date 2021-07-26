package main

func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }
    var fn func(r1, r2 *TreeNode) bool
    fn = func(r1, r2 *TreeNode) bool {
        if r1 == nil && r2 == nil {
            return true
        }
        if r1 == nil || r2 == nil || r1.Val != r2.Val {
            return false
        }
        return fn(r1.Left, r2.Right) && fn(r1.Right, r2.Left)
    }
    return fn(root.Left, root.Right)
}
