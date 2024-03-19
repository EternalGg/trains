package deleteMiddle

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow, fast, last := head, head, head
	for fast != nil {
		if fast.Next == nil {
			last.Next = slow.Next
			break
		}
		fast = fast.Next
		if fast.Next == nil {
			slow.Next = slow.Next.Next
			break
		}
		fast = fast.Next
		last = slow
		slow = slow.Next
	}
	return head
}
