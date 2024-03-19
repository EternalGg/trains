package removeZeroSumSublists

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeZeroSumSublists(head *ListNode) *ListNode {
	nList := make([]*ListNode, 0)
	for head != nil {
		nList = append(nList, head)
		head = head.Next
	}

	return head
}
