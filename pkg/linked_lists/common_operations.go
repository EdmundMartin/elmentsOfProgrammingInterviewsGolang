package linked_lists

func InsertAfter(node *ListNode, newNode *ListNode) {
	newNode.Next = node.Next
	node.Next = newNode
}

func SearchList(l *ListNode, value int) *ListNode {
	for l != nil && l.Data != value {
		l = l.Next
	}
	// If value is not in list - L will be nil
	return l
}

func DeleteFromList(node *ListNode) {
	node.Next = node.Next.Next
}

func DeletionFromList(nodeForDeletion *ListNode) {
	nodeForDeletion.Data = nodeForDeletion.Next.Data
	nodeForDeletion.Next = nodeForDeletion.Next.Next
}
