package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyNode := &ListNode{Val: 0}
	cur := dummyNode
	n1, n2, flag := 0, 0, 0
	for l1 != nil || l2 != nil || flag != 0 {
		if l1 == nil {
			n1 = 0
		} else {
			n1 = l1.Val
			l1 = l1.Next
		}

		if l2 == nil {
			n2 = 0
		} else {
			n2 = l2.Val
			l2 = l2.Next
		}

		cur.Next = &ListNode{Val: (n1 + n2 + flag) % 10}
		cur = cur.Next
		flag = (n1 + n2 + flag) / 10
	}
	return dummyNode.Next
}
