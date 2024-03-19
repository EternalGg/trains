package myqueue

type MyQueue struct {
	list   []int
	length int
}

func Constructor() MyQueue {
	q := new(MyQueue)
	q.list = make([]int, 0)
	q.length = 0
	return *q
}

func (this *MyQueue) Push(x int) {
	this.list = append(this.list, x)
	this.length++
}

func (this *MyQueue) Pop() int {
	if this.length == 0 {
		return -1
	} else {
		result := this.list[0]
		this.list = this.list[1:]
		this.length--
		return result
	}
}

func (this *MyQueue) Peek() int {
	if this.length == 0 {
		return -1
	} else {
		return this.list[0]
	}
}

func (this *MyQueue) Empty() bool {
	if this.length == 0 {
		return true
	} else {
		return false
	}
}
