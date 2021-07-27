package main

func reorderList(head *ListNode) {
    nodes := make([]*ListNode, 0)
    cur := head
    for cur != nil {
        node := cur
        cur = cur.Next
        node.Next = nil
        nodes = append(nodes, node)
    }
    dummyNode := &ListNode{}
    cur = dummyNode
    nodes1 := make([]*ListNode, len(nodes)/2)
    nodes2 := make([]*ListNode, len(nodes)-len(nodes)/2)
    copy(nodes1, nodes[:len(nodes)/2])
    copy(nodes2, nodes[len(nodes)/2:])
    //reverseListArray(nodes1)
    reverseListArray(nodes2)
    for i := 0; i < len(nodes)/2+1; i++ {
        if i < len(nodes1) {
            cur.Next = nodes1[i]
            cur = cur.Next
        }
        if i < len(nodes2) {
            cur.Next = nodes2[i]
            cur = cur.Next
        }
    }
    printList(dummyNode.Next)
}

func reverseListArray(nodes []*ListNode) {
    i := 0
    j := len(nodes) - 1
    for i < j {
        nodes[i], nodes[j] = nodes[j], nodes[i]
        i++
        j--
    }
}
