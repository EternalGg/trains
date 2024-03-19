package nodesBetweenCriticalPoints

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func nodesBetweenCriticalPoints(head *ListNode) []int {
	first, last, key, min := -1, -1, 1, 0

	for head != nil {
		key++
		if head.Next.Next == nil {
			break
		}
		if (head.Val > head.Next.Val && head.Next.Next.Val > head.Next.Val) || (head.Val < head.Next.Val && head.Next.Val < head.Next.Next.Val) {
			if first == -1 {
				first = key
			} else if last == -1 {
				last = key

				min = last - first
				fmt.Println(min)
			} else {
				if key-last < min {
					min = key - last
				}
				last = key
			}
		}
		head = head.Next
	}
	result := make([]int, 2)
	if last == -1 || first == -1 {
		result[0] = -1
		result[1] = -1
	} else {
		result[0] = min
		result[1] = last - first
	}
	return result
}
