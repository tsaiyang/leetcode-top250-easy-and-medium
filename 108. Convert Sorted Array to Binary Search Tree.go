package main

func sortedArrayToBST(nums []int) *TreeNode {
    var fn func(start, end int) *TreeNode
    fn = func(start, end int) *TreeNode {
        if start > end {
            return nil
        }
        mid := (start + end) / 2
        root := &TreeNode{Val: nums[mid]}
        root.Left = fn(start, mid-1)
        root.Right = fn(mid+1, end)
        return root
    }
    return fn(0, len(nums)-1)
}

// []int{-10, -3, 0, 5, 9}
func findMin(nums []int, start, end int) int {
    result := start
    for i := start; i <= end; i++ {
        if nums[i] < nums[result] {
            result = i
        }
    }
    return result
}
