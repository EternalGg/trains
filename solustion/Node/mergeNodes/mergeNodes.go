package mergeNodes

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeNodes(head *ListNode) *ListNode {
	list, resultList := make([]int, 0), make([]int, 0)
	for head != nil {
		list = append(list, head.Val)
		head = head.Next
	}
	add, left, right, start := 0, 0, 0, false
	for key, value := range list {
		add += value
		right++
		if value == 0 {
			if start {
				resultList = append(resultList, add)
			} else {
				start = true
				resultList = append(resultList, list[0:right]...)
			}
			left = key
			add = 0
		}
	}
	if left != right {
		resultList = append(resultList, list[left:right]...)
	}
	if resultList[0] == 0 {
		resultList = resultList[1:]
	}
	if resultList[len(resultList)-1] == 0 {
		resultList = resultList[:len(resultList)-2]
	}
	fmt.Println(resultList)
	return listToNode(resultList)
}

func listToNode(nums []int) *ListNode {
	head := new(ListNode)
	node := head
	for key, num := range nums {
		node.Val = num
		if key == len(nums) {
			break
		}
		newNode := new(ListNode)
		node.Next = newNode
		node = newNode
	}
	return head
}
