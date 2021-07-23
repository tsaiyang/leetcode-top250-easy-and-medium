package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || n < 1 {
		return head
	}
	cur := head
	for cur != nil {
		n--
		cur = cur.Next
	}
	if n == 0 {
		return head.Next
	}
	if n < 0 {
		cur = head
		for n < -1 {
			cur = cur.Next
		}
		cur.Next = cur.Next.Next
	}
	return head
}
