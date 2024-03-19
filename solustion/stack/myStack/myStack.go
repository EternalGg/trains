package myStack

type MyStack struct {
	list   []int
	length int
}

func Constructor() MyStack {
	s := new(MyStack)
	s.list = make([]int, 0)
	s.length = 0
	return *s
}

func (this *MyStack) Push(x int) {
	this.list = append(this.list, x)
	this.length++
}

func (this *MyStack) Pop() int {
	if this.length == 0 {
		return -1
	} else {
		result := this.list[this.length-1]
		this.list = this.list[:this.length-1]
		this.length--
		return result
	}
}

func (this *MyStack) Top() int {
	if this.length == 0 {
		return -1
	} else {
		return this.list[this.length-1]
	}
}

func (this *MyStack) Empty() bool {
	if this.length == 0 {
		return true
	} else {
		return false
	}
}
