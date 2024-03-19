package GetRandom

import (
	"math/rand"
	"time"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct {
	List   []*ListNode
	Length int
}

func Constructor(head *ListNode) Solution {
	s := new(Solution)
	s.List = make([]*ListNode, 0)
	for head != nil {
		s.List = append(s.List, head)
		s.Length++
		head = head.Next
	}
	return *s
}

func (this *Solution) GetRandom() int {
	if this.Length == 1 {
		return this.List[0].Val
	}
	rander := rand.New(rand.NewSource(time.Now().UnixNano()))
	n1 := rander.Intn((this.Length) - 1)
	//data := rand.Int63n()
	return this.List[n1].Val
}
