package main

import "sort"

func reorderList1(head *ListNode) {
    nodesArray := make([]*ListNode, 0)
    cur := head
    for cur != nil {
        node := cur
        cur = cur.Next
        node.Next = nil
        nodesArray = append(nodesArray, node)
    }
    sort.Sort(ListNodeSlice(nodesArray))
    dummyNode := &ListNode{}
    curNode := dummyNode
    for _, node := range nodesArray {
        curNode.Next = node
        curNode = curNode.Next
    }
    head = dummyNode.Next
}

type ListNodeSlice []*ListNode

func (l ListNodeSlice) Len() int {
    return len(l)
}

func (l ListNodeSlice) Less(i, j int) bool {
    return l[i].Val < l[j].Val
}

func (l ListNodeSlice) Swap(i, j int) {
    l[i], l[j] = l[j], l[i]
}
