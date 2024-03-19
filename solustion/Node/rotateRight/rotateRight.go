package rotateRight

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	list := make([]*ListNode, 0)
	for head != nil {
		list = append(list, head)
		head = head.Next
	}
	if k >= len(list) {
		k = k % len(list)
	}
	newlist := append(list[len(list)-k:], list[:len(list)-k]...)
	for i := 0; i < len(newlist)-1; i++ {
		newlist[i].Next = newlist[i+1]
	}
	newlist[len(newlist)-1].Next = nil
	return newlist[0]
}
