package main

func buildTree1(inorder []int, postorder []int) *TreeNode {
    var fn func(inStart, inEnd, postStart, postEnd int) *TreeNode
    fn = func(inStart, inEnd, postStart, postEnd int) *TreeNode {
        if postEnd < postStart {
            return nil
        }
        root := &TreeNode{Val: postorder[postEnd]}
        count := 0
        for i := inStart; i <= inEnd; i++ {
            if inorder[i] == postorder[postEnd] {
                break
            }
            count++
        }
        root.Left = fn(inStart, inStart+count, postStart, postStart+count)
        root.Right = fn(inStart+count+2, inEnd, postStart+count+1, postEnd-1)
        return root
    }
    return fn(0, len(inorder)-1, 0, len(postorder)-1)
}
