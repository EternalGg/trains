package mergeInBetween

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	beforeNode, l1key, result := new(ListNode), 0, list1
	for l1key != a-1 {
		list1 = list1.Next
		l1key++
	}
	beforeNode = list1

	nextNode := new(ListNode)
	for l1key != b+1 {
		list1 = list1.Next
		l1key++
	}
	nextNode = list1

	beforeNode.Next = list2
	for list2.Next != nil {
		list2 = list2.Next
	}
	list2.Next = nextNode
	return result
}
