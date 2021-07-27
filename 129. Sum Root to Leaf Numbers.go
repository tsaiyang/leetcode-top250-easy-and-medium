package main

func sumNumbers(root *TreeNode) int {
    result := 0
    var arr []int
    var fn func(root *TreeNode)
    fn = func(root *TreeNode) {
        if root == nil {
            return
        }
        if root.Left == nil && root.Right == nil {
            var sum int
            for i := 0; i < len(arr); i++ {
                sum = sum*10 + arr[i]
            }
            sum = sum*10 + root.Val
            result += sum
            return
        }
        arr = append(arr, root.Val)
        if root.Left != nil {
            fn(root.Left)
        }
        if root.Right != nil {
            fn(root.Right)
        }
        arr = arr[:len(arr)-1]
    }
    fn(root)
    return result
}
