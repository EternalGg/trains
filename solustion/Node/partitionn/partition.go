package partitionn

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	smaller, bigger := make([]*ListNode, 0), make([]*ListNode, 0)
	for head != nil {
		if head.Val >= x {
			bigger = append(bigger, head)
		} else {
			smaller = append(smaller, head)
		}
		head = head.Next
	}
	list := append(smaller, bigger...)
	for i := 0; i < len(list)-1; i++ {
		list[i].Next = list[i+1]
	}
	list[len(list)-1].Next = nil
	return list[0]
}
