package linked_lists

func RemoveKthNode(l *ListNode, k int) *ListNode {
	dummyHead := &ListNode{
		Data: 0,
		Next: l,
	}
	first := dummyHead.Next
	for i := 0; i < k; i++ {
		first = first.Next
	}
	second := dummyHead
	for first != nil {
		first, second = first.Next, second.Next
	}
	second.Next = second.Next.Next
	return dummyHead.Next
}
