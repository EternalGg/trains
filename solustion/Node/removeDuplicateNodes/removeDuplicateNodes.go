package removeDuplicateNodes

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeDuplicateNodes(head *ListNode) *ListNode {
	hMap, list := make(map[int]bool), make([]*ListNode, 0)
	for head != nil {
		if !hMap[head.Val] {
			list = append(list, head)
			hMap[head.Val] = true
		}
		head = head.Next
	}
	for i := 0; i < len(list)-1; i++ {
		list[i].Next = list[i+1]
	}
	list[len(list)-1].Next = nil
	return list[0]
}
