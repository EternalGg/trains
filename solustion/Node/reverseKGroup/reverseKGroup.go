package reverseKGroup

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	nList := make([]*ListNode, 0)
	for head != nil {
		nList = append(nList, head)
		head = head.Next
	}
	reversrList := make([]*ListNode, 0)

	for i := 1; i <= len(nList)/k; i++ {
		now := (i * k) - 1
		for j := now; j > now-2; j-- {
			reversrList = append(reversrList, nList[j])
		}
	}
	fmt.Println(len(reversrList), nList[len(reversrList):])
	reversrList = append(reversrList, nList[len(reversrList):]...)
	for i := 0; i < len(reversrList)-1; i++ {
		reversrList[i].Next = reversrList[i+1]
		fmt.Println(reversrList[i].Val)
	}
	reversrList[len(reversrList)-1].Next = nil
	return nList[0]
}
