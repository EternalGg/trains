package queue

type CQueue struct {
	list   []int
	length int
}

func Constructor() CQueue {
	q := new(CQueue)
	q.list = make([]int, 0)
	q.length = 0
	return *q
}

func (this *CQueue) AppendTail(value int) {
	this.list = append(this.list, value)
	this.length++
}

func (this *CQueue) DeleteHead() int {
	if this.length == 0 {
		return -1
	}
	result := this.list[0]
	this.list = this.list[1:]
	this.length--
	return result
}
