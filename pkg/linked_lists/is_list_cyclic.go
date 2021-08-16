package linked_lists


func cycleLength(end *ListNode) int {
	start, step := end, 0
	for {
		step++
		start = start.Next
		if start == end {
			return step
		}
	}
}

func HasCycle(head *ListNode) *ListNode {
	fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
		if slow == fast {
			cycleLenAdvanced := head
			cycleLength := cycleLength(slow)
			for i := 0; i < cycleLength; i++ {
				cycleLenAdvanced = cycleLenAdvanced.Next
			}
			it := head
			for it != cycleLenAdvanced {
				it = it.Next
				cycleLenAdvanced = cycleLenAdvanced.Next
			}
			return it
		}
	}
	return nil
}
