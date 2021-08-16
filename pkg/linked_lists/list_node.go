package linked_lists

import (
	"bytes"
	"fmt"
)

type ListNode struct {
	Data int
	Next *ListNode
}

func (l *ListNode) String() string {
	var out bytes.Buffer
	node := l
	for node != nil {
		out.WriteString(fmt.Sprintf("%d, ", node.Data))
		node = node.Next
	}
	return out.String()
}