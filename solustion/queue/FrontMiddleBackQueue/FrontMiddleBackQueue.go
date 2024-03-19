package FrontMiddleBackQueue

import "fmt"

type FrontMiddleBackQueue struct {
	List   []int
	length int
}

func (f *FrontMiddleBackQueue) isEmpty() bool {
	if f.length == 0 {
		return true
	} else {
		return false
	}
}

func Constructor() FrontMiddleBackQueue {
	f := new(FrontMiddleBackQueue)
	f.List = make([]int, 0)
	f.length = 0
	return *f
}

func (this *FrontMiddleBackQueue) PushFront(val int) {
	front := make([]int, 1)
	front[0] = val
	front = append(front, this.List...)
	this.length++
	this.List = front
}

func (this *FrontMiddleBackQueue) PushMiddle(val int) {
	half := len(this.List) / 2
	head := this.List[:half]
	rear := this.List[half:]
	last := make([]int, len(rear))
	for i := 0; i < len(last); i++ {
		last[i] = rear[i]
	}
	newOne := append(head, val)
	newOne = append(newOne, last...)
	fmt.Println(this.List, head, last)
	this.List = newOne
	fmt.Println(this.List, head, last)
	this.length++
}

func (this *FrontMiddleBackQueue) PushBack(val int) {
	this.List = append(this.List, val)
	this.length++
}

func (this *FrontMiddleBackQueue) PopFront() int {
	if this.isEmpty() {
		return -1
	} else {
		result := this.List[0]
		this.List = this.List[1:]
		this.length--
		return result
	}
}

func (this *FrontMiddleBackQueue) PopMiddle() int {
	if this.isEmpty() {
		return -1
	} else {
		half := this.length / 2
		result := this.List[half]
		head := this.List[:half]
		rear := this.List[half+1:]
		this.List = append(head, rear...)
		fmt.Println(this.List)
		this.length--
		return result
	}
}

func (this *FrontMiddleBackQueue) PopBack() int {
	if this.isEmpty() {
		return -1
	} else {
		result := this.List[this.length-1]
		this.List = this.List[:this.length-1]
		this.length--
		return result
	}
}
