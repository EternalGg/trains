package reverseBetween

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseBetween(head *ListNode, left int, right int) *ListNode {
	if right-left == 0 {
		return head
	}
	left -= 1
	right -= 1
	list := make([]*ListNode, 0)
	for head != nil {
		list = append(list, head)
		head = head.Next
	}
	key := 0
	for key != right {
		if key > left && key <= right {
			fmt.Println("helllo")
			list[key].Next = list[key-1]
		}
		key++
	}
	if left == 0 {
		if len(list) > right+1 {
			list[left].Next = list[right+1]
		} else {
			list[left].Next = nil
		}
		return list[right]
	}
	if len(list) <= right+1 {
		list[left].Next = nil
	} else {
		list[left].Next = list[right+1]
	}

	list[left-1].Next = list[right]
	return list[0]
}
