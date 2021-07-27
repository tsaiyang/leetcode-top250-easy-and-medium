package main

func postorderTraversal(root *TreeNode) []int {
    var result []int
    var fn func(root *TreeNode)
    fn = func(root *TreeNode) {
        if root == nil {
            return
        }
        fn(root.Left)
        fn(root.Right)
        result = append(result, root.Val)
    }
    fn(root)
    return result
}
