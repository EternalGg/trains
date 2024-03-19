package splitListToParts

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func splitListToParts(head *ListNode, k int) []*ListNode {
	list := make([]*ListNode, 0)
	for head != nil {
		list = append(list, head)
		head = head.Next
	}
	avg, result := len(list)/k, make([]*ListNode, k)

	if avg == 0 {
		for i := 0; i < len(list); i++ {
			list[i].Next = nil
			result[i] = list[i]
		}
	} else {
		more := len(list) % avg
		last := 0
		fmt.Println(more, len(list), avg)
		for i := 0; i < k; i++ {
			header := last
			last += avg
			if more > 0 {
				last += 1
				more--
			}
			fmt.Println(more, last, header)
			now := list[header:last]
			now[len(now)-1].Next = nil
			result[i] = now[0]
		}
	}

	return result
}
