package main

func pathSum(root *TreeNode, targetSum int) [][]int {
    var result [][]int
    var arr []int
    var fn func(root *TreeNode)
    var sum int
    fn = func(root *TreeNode) {
        if root == nil {
            return
        }
        arr = append(arr, root.Val)
        sum += root.Val
        if root.Left == nil && root.Right == nil {
            if sum == targetSum {
                tmp := make([]int, len(arr))
                copy(tmp, arr)
                result = append(result, tmp)
            }
            arr = arr[:len(arr)-1]
            return
        }
        fn(root.Left)
        fn(root.Right)
        arr = arr[:len(arr)-1]
        sum -= root.Val
    }
    fn(root)
    return result
}
