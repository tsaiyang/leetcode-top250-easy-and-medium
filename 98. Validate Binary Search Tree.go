package main

import "math"

func isValidBST(root *TreeNode) bool {
    var fn func(root *TreeNode, maxEdge, minEdge int64) bool
    fn = func(root *TreeNode, maxEdge, minEdge int64) bool {
        if root == nil {
            return true
        }
        if int64(root.Val) < minEdge || int64(root.Val) > maxEdge {
            return false
        }
        return fn(root.Left, int64(root.Val), minEdge) && fn(root.Right, maxEdge, int64(root.Val))
    }
    return fn(root, math.MaxInt64, math.MinInt64)
}
