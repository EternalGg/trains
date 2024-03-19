package MyCircularQueue

type MyCircularQueue struct {
	List []int
	max  int
}

func Constructor(k int) MyCircularQueue {
	m := new(MyCircularQueue)
	m.List = make([]int, 0)
	m.max = k
	return *m
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	} else {
		this.List = append(this.List, value)
		return true
	}
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	} else {
		this.List = this.List[1:]
		return true
	}
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	} else {
		return this.List[0]
	}
}

func (this *MyCircularQueue) Rear() int {
	if !this.IsEmpty() {
		return this.List[len(this.List)-1]
	} else {
		return -1
	}
}

func (this *MyCircularQueue) IsEmpty() bool {
	if len(this.List) == 0 {
		return true
	} else {
		return false
	}
}

func (this *MyCircularQueue) IsFull() bool {
	if len(this.List) == this.max {
		return true
	} else {
		return false
	}
}
