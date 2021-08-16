package linked_lists


func MergeTwoSortedLists(first, second *ListNode) *ListNode {
	dummyHead := &ListNode{}
	tail := dummyHead

	for first != nil && second != nil {
		if first.Data <= second.Data {
			tail.Next, first = first, first.Next
		} else {
			tail.Next, second = second, second.Next
		}
		tail = tail.Next
	}
	if first != nil {
		tail.Next = first
	} else {
		tail.Next = second
	}
	return dummyHead.Next
}
