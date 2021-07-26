package main

func sortedListToBST(head *ListNode) *TreeNode {
    if head == nil {
        return nil
    }
    var nodes []*ListNode
    for head != nil {
        nodes = append(nodes, head)
        head = head.Next
    }

    var fn func(start, end int) *TreeNode
    fn = func(start, end int) *TreeNode {
        if start > end {
            return nil
        }
        mid := (start + end) / 2
        root := &TreeNode{Val: nodes[mid].Val}
        root.Left = fn(start, mid-1)
        root.Right = fn(mid+1, end)
        return root
    }
    return fn(0, len(nodes)-1)
}
