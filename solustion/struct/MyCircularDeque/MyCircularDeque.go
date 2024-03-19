package MyCircularDeque

import "fmt"

type listNode struct {
	Val  int
	Next *listNode
	Head *listNode
}

type MyCircularDeque struct {
	Head   *listNode
	Last   *listNode
	Max    int
	length int
}

func Constructor(k int) MyCircularDeque {
	m := new(MyCircularDeque)
	m.Max = k
	m.length = 0
	return *m
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	} else {
		node := new(listNode)
		node.Val = value
		if this.IsEmpty() {
			this.Head = node
			this.Last = node
		} else {
			node.Next = this.Head
			this.Head.Head = node
			this.Head = node
		}
		this.length++
		return true
	}
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	} else {
		node := new(listNode)
		node.Val = value
		fmt.Println(this.length)
		if this.IsEmpty() {
			this.Head = node
			this.Last = node
		} else {
			node.Head = this.Last
			this.Last.Next = node
			this.Last = node
		}
		this.length++
		return true
	}
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	} else {
		if this.length == 1 {
			this.Head = nil
			this.Last = nil
		} else {
			this.Head = this.Head.Next
			this.Head.Head = nil
		}
		this.length--
		return true
	}
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	} else {
		if this.length == 1 {
			this.Head = nil
			this.Last = nil
		} else {
			this.Last = this.Last.Head
			this.Last.Next = nil
		}
		this.length--
		return true
	}
}

func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	} else {
		return this.Head.Val
	}
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	} else {
		return this.Last.Val
	}
}

func (this *MyCircularDeque) IsEmpty() bool {
	if this.length == 0 {
		return true
	} else {
		return false
	}
}

func (this *MyCircularDeque) IsFull() bool {
	if this.length == this.Max {
		return true
	} else {
		return false
	}
}
