package main

type RNode struct {
    Val    int
    Next   *RNode
    Random *RNode
}

func copyRandomList(head *RNode) *RNode {
    nodeMap := make(map[*RNode]*RNode)
    cur := head
    for cur != nil {
        nodeMap[cur] = &RNode{Val: cur.Val}
        cur = cur.Next
    }
    cur = head
    for cur != nil {
        nodeMap[cur].Random = nodeMap[cur.Random]
        nodeMap[cur].Next = nodeMap[cur.Next]
    }
    return nodeMap[head]
}
