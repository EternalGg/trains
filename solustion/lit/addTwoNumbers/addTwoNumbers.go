package addTwoNumbers

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type struck struct {
	nList  []int
	length int
}

func (s *struck) pop() int {
	if s.length == 0 {
		return 0
	}
	re := s.nList[s.length-1]
	s.nList = s.nList[:s.length-1]
	s.length = s.length - 1
	return re
}

func (s *struck) push(num int) {
	s.nList = append(s.nList, num)
	s.length = s.length + 1
	fmt.Println(s.length)
	return
}

func (s *struck) stackInit() {
	s.nList = make([]int, 0)
	s.length = 0
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head, l1S, l2S, up := new(ListNode), new(struck), new(struck), false
	l1S.stackInit()
	l2S.stackInit()
	for l1 != nil {
		l1S.push(l1.Val)
		fmt.Println(l1S.nList)
		l1 = l1.Next
	}
	for l2 != nil {
		l2S.push(l2.Val)
		l2 = l2.Next
	}
	fmt.Println(l1S.nList, l2S.nList)
	for l1S.length != 0 && l2S.length != 0 {

		n1, n2 := l1S.pop(), l2S.pop()
		val := n1 + n2
		if up {
			val += 1
		}
		val = val % 10
		if val >= 10 {
			up = true
		} else {
			up = false
		}
		head.Val = val
		fmt.Println(val)
		newOne := new(ListNode)
		newOne.Next = head
		head = newOne
	}
	if up {
		head.Val = 1
	}
	return head
}
