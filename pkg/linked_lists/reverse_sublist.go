package linked_lists

func ReverseSublist(l *ListNode, start, end int) *ListNode {
	dummyHead := &ListNode{Data: 0, Next: l}
	subListHead := dummyHead
	for i := 0; i < start; i++ {
		subListHead = subListHead.Next
	}

	subListIter := subListHead.Next
	for i := 0; i < end-start; i++ {
		temp := subListIter.Next
		subListIter.Next, temp.Next, subListHead.Next = temp.Next, subListHead.Next, temp
	}
	return dummyHead.Next
}
