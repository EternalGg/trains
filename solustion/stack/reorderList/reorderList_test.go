package reorderList

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	list := make([]*ListNode, 0)
	for head != nil {
		list = append(list, head)
		head = head.Next
	}

	half := len(list) / 2

	newList := make([]*ListNode, 0)
	for i := 0; i < half; i++ {
		newList = append(newList, list[i])
		newList = append(newList, list[len(list)-i-1])
	}
	if len(list)%2 != 0 {
		newList = append(newList, list[half])
	}
	for i := 0; i < len(newList)-1; i++ {
		newList[i].Next = newList[i+1]
	}
	newList[len(newList)-1].Next = nil
}
