package linked_lists

func lengthList(l *ListNode) int {
	length := 0
	for l != nil {
		length++
		l = l.Next
	}
	return length
}

func OverlappingTerminalLists(first, second *ListNode) bool {
	firstLength, secondLength := lengthList(first), lengthList(second)
	if firstLength > secondLength {
		first, second = second, first
	}

	for i := 0; i < lengthList(second)-lengthList(first); i++ {
		first = first.Next
	}

	for first != nil && second != nil && first != second {
		first, second = first.Next, second.Next
	}
	return first != nil
}
