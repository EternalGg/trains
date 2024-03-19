package doubleIt

type ListNode struct {
	Val  int
	Next *ListNode
}

func doubleIt(head *ListNode) *ListNode {
	nodeList := []*ListNode{}
	for head != nil {
		nodeList = append(nodeList, head)
		head = head.Next
	}
	length, b := len(nodeList), false
	for i := length - 1; i >= 0; i-- {
		now := nodeList[i].Val * 2
		if b {
			now++
		}
		if now >= 10 {
			b = true
			nodeList[i].Val = now % 10
		} else {
			nodeList[i].Val = now
			b = false
		}
	}
	if b {
		return &ListNode{Val: 1, Next: nodeList[0]}
	}
	return nodeList[0]
}
