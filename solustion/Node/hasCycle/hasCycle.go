package hasCycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	nodeMap := make(map[*ListNode]bool)

	for nodeMap[head] == false {
		if head == nil {
			return false
		}
		nodeMap[head] = true
		head = head.Next
	}

	return true
}
