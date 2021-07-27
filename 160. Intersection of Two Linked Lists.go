package main

func getIntersectionNode(headA, headB *ListNode) *ListNode {
    aMap := make(map[*ListNode]bool)
    cur := headA
    for cur != nil {
        aMap[cur] = true
        cur = cur.Next
    }

    cur = headB
    for cur != nil {
        if aMap[cur] {
            return cur
        }
        cur = cur.Next
    }
    return nil
}

func getIntersectionNode1(headA, headB *ListNode) *ListNode {
    curA := headA
    curB := headB
    countA := 0
    countB := 0

    for curA != nil {
        countA++
        curA = curA.Next
    }
    for curB != nil {
        countB++
        curB = curB.Next
    }

    curA = headA
    curB = headB
    if countA > countB {
        for countA > countB {
            curA = curA.Next
            countA--
        }
    } else {
        for countB > countA {
            curB = curB.Next
            countB--
        }
    }
    for curA != nil && curB != nil {
        if curA == curB {
            return curA
        }
        curA = curA.Next
        curB = curB.Next
    }
    return nil
}
