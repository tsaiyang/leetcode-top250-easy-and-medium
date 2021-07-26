package main

func buildTree(preorder []int, inorder []int) *TreeNode {
    var fn func(preStart, preEnd, inStart, inEnd int) *TreeNode
    fn = func(preStart, preEnd, inStart, inEnd int) *TreeNode {
        if preStart > preEnd {
            return nil
        }
        root := &TreeNode{Val: preorder[preStart]}
        count := 0
        for i := inStart; i <= inEnd; i++ {
            if inorder[i] == preorder[preStart] {
                break
            }
            count++
        }
        root.Left = fn(preStart+1, preStart+count, inStart, inStart+count-1)
        root.Right = fn(preStart+count+1, preEnd, inStart+count+1, inEnd)
        return root
    }
    return fn(0, len(preorder)-1, 0, len(inorder)-1)
}
