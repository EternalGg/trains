package getIntersectionNode

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	aMap := make(map[*ListNode]*ListNode)
	for headA != nil {
		aMap[headA] = headA
		headA = headA.Next
	}
	for aMap[headB] == nil {
		if headB == nil {
			return nil
		}
		headB = headB.Next

	}
	return headB
}
