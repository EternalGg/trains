package removeNodes

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNodes(head *ListNode) *ListNode {
	nList := make([]*ListNode, 0)
	node := head
	for head != nil {
		nList = append(nList, head)
		head = head.Next
	}
	last := 0
	index := make([]*ListNode, 0)
	for i := len(nList) - 1; i >= 0; i-- {
		if nList[i].Val > last {
			index = append(index, nList[i])
			last = nList[i].Val
		}
	}
	if len(index) == 1 {
		return node
	}
	for i := len(index) - 1; i > 0; i-- {
		index[i].Next = index[i-1]
	}
	index[0].Next = nil
	return index[len(index)-1]
}
