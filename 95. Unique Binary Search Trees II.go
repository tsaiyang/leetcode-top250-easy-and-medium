package main

func generateTrees(n int) []*TreeNode {
    var result []*TreeNode
    var fn func(start, end int) []*TreeNode
    fn = func(start, end int) []*TreeNode {
        if start > end {
            return []*TreeNode{nil}
        }
        for i := start; i <= end; i++ {
            root := &TreeNode{Val: i}
            lefts := fn(start, i-1)
            rights := fn(i+1, end)
            for i := range lefts {
                for j := range rights {
                    root.Left = lefts[i]
                    root.Right = rights[j]
                    result = append(result, copyTree(root))
                }
            }
        }
        return result
    }
    return fn(1, n)
}

func copyTree(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    newRoot := &TreeNode{Val: root.Val}
    newRoot.Left = copyTree(root.Left)
    newRoot.Right = copyTree(root.Right)
    return newRoot
}
