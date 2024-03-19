package getKthFromEnd

type ListNode struct {
	Val  int
	Next *ListNode
}

type NodeQueue struct {
	list   []*ListNode
	length int
	max    int
}

func NodeQueueInit(length int) *NodeQueue {
	q := new(NodeQueue)
	q.max = length
	q.list = make([]*ListNode, q.max)
	q.length = 0
	return q
}

func (nq *NodeQueue) push(node *ListNode) {
	if nq.length < nq.max {
		nq.list[nq.length] = node
		nq.length++
	} else {
		nq.list = nq.list[1:]
		nq.list = append(nq.list, node)
	}
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	nq := NodeQueueInit(k)
	for head != nil {
		nq.push(head)
		head = head.Next
	}
	return nq.list[0]
}
