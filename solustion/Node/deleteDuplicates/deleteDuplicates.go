package deleteDuplicates

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func DeleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	node, remainList, countMap, last, firstNode := head, make([]int, 0), make(map[int]int), 0, make(map[int]*ListNode)

	for node != nil {
		if countMap[node.Val] >= 1 {
			last = node.Val
			for node.Val == last {
				if node.Next == nil {
					break
				}
				node = node.Next
			}
			length := len(remainList)
			if length == 1 {
				remainList = make([]int, 0)
			} else {
				remainList = remainList[0 : length-1]
			}

		}
		countMap[node.Val] = 1
		remainList = append(remainList, node.Val)
		firstNode[node.Val] = node
		node = node.Next
	}

	fmt.Println(remainList)
	if len(remainList) == 0 {
		return nil
	}
	firstHead := new(ListNode)
	for key, value := range remainList {
		if key == 0 {
			firstHead = firstNode[value]
		}
		if key == len(remainList)-1 {
			break
		}
		firstNode[value].Next = firstNode[remainList[key+1]]
	}
	return firstHead
}
