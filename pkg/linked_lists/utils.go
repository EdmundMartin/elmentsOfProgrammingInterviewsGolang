package linked_lists

func SingleLinkedListFromSlice(contents []int) *ListNode {
	head := &ListNode{Data: contents[0]}
	node := head
	for _, val := range contents[1:] {
		currentNode := &ListNode{Data: val}
		node.Next = currentNode
		node = currentNode
	}
	return head
}