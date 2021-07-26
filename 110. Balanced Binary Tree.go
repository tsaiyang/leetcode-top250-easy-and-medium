package main

func isBalanced(root *TreeNode) bool {
    var fn func(root *TreeNode) (int, bool)
    fn = func(root *TreeNode) (int, bool) {
        if root == nil {
            return 0, true
        }
        leftDepth, leftBalance := fn(root.Left)
        rightDepth, rightBalance := fn(root.Right)
        if !leftBalance || !rightBalance || abs(leftDepth-rightDepth) > 2 {
            return max(leftDepth, rightDepth) + 1, false
        }
        return max(leftDepth, rightDepth) + 1, true
    }
    _, balance := fn(root)
    return balance
}
