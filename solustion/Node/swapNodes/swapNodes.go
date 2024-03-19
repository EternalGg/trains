package swapNodes

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapNodes(head *ListNode, k int) *ListNode {
	list := make([]*ListNode, 0)
	for head != nil {
		list = append(list, head)
		head = head.Next
	}
	middle := list[len(list)-k]
	list[len(list)-k] = list[k-1]
	list[k-1] = middle

	for i := 0; i < len(list)-1; i++ {
		list[i].Next = list[i+1]
	}
	list[len(list)-1].Next = nil
	return list[0]
}
