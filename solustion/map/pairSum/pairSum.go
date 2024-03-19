package pairSum

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	pairMap, key := make(map[int]int), 0
	for head != nil {
		pairMap[key] = head.Val
		head = head.Next
	}
	max := 0
	for i := 0; i < key; i++ {
		fmt.Println(pairMap[key-1-i] + pairMap[i])
		if pairMap[key-1-i]+pairMap[i] > max {
			max = pairMap[key-1-i] + pairMap[i]
		}
	}
	return max
}
