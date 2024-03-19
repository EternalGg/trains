package addTwoNumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	last, head := 0, new(ListNode)
	node := head

	for l1 == nil && l2 == nil {
		now, up := last, false
		switch {
		case l1 == nil:
			now += l2.Val
			l2 = l2.Next
		case l2 == nil:
			now += l1.Val
			l1 = l1.Next
		default:
			now += l1.Val + l2.Val
			l2 = l2.Next
			l1 = l1.Next
		}
		node.Val, up = TenUp(now)
		if up {
			last = 1
		} else {
			last = 0
		}
		newNode := new(ListNode)
		node.Next = newNode
		node = node.Next
	}
	if last == 0 {
		node = nil
	} else {
		node.Val = 1
	}
	return head
}

func TenUp(num int) (single int, isUp bool) {
	if num > 10 {
		return num % 10, true
	} else {
		return num, false
	}
}
