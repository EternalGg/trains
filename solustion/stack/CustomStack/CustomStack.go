package CustomStack

import "math"

type CustomStack struct {
	list   []int
	max    int
	length int
}

func (C *CustomStack) isZero() bool {
	if C.length == 0 {
		return true
	} else {
		return false
	}
}

func (C *CustomStack) isFull() bool {
	if C.length == C.max {
		return true
	} else {
		return false
	}
}

func Constructor(maxSize int) CustomStack {
	c := new(CustomStack)
	c.list = make([]int, maxSize)
	c.max = maxSize
	c.length = 0
	return *c
}

func (this *CustomStack) Push(x int) {
	if this.isFull() {
		return
	} else {
		this.list[this.length] = x
		this.length++
	}
}

func (this *CustomStack) Pop() int {
	if this.isZero() {
		return -1
	} else {
		result := this.list[this.length-1]
		//this.list[this.max-1] = 0
		this.length -= 1
		return result
	}
}

func (this *CustomStack) Increment(k int, val int) {
	t := math.Min(float64(k), float64(this.length))
	for i := 0; i < int(t); i++ {
		this.list[i] += val
	}
}
