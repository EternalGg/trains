package insertGreatestCommonDivisors

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	node := head
	for node.Next != nil {
		if node.Val == node.Next.Val {
			newNode := ListNode{Val: node.Val, Next: node.Next}
			newNode.Next = node.Next
			node.Next = &newNode
			node = node.Next.Next
		} else {
			biggest, insert := Max(node.Next.Val, node.Val), 1
			for i := biggest; i >= 1; i-- {
				if node.Val%i == 0 && node.Next.Val%i == 0 {
					insert = i
					break
				}
			}
			newNode := ListNode{Val: insert, Next: node.Next}
			newNode.Next = node.Next
			node.Next = &newNode
			node = node.Next.Next
		}
	}
	return head
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
