package main

// leaf to leaf max
// any node to any node max
// root to any node max
// root to leaf max
func hasPathSum(root *TreeNode, targetSum int) bool {
    preSum := 0
    var fn func(root *TreeNode) bool
    fn = func(root *TreeNode) bool {
        if root == nil {
            return false
        }
        sum := preSum + root.Val
        if root.Left == nil && root.Right == nil {
            if sum == targetSum {
                return true
            }
            return false
        }
        preSum += root.Val
        flag := fn(root.Left) || fn(root.Right)
        preSum -= root.Val
        return flag
    }
    return fn(root)
}
