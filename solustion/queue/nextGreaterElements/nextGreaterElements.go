package nextGreaterElements

import "fmt"

type postion struct {
	value int
	key   int
}

type MyCircularQueue struct {
	List []postion
	max  int
}

func Constructor(k int) MyCircularQueue {
	m := new(MyCircularQueue)
	m.List = make([]postion, 0)
	m.max = k
	return *m
}

func (this *MyCircularQueue) EnQueue(value, key int) bool {
	if this.IsFull() {
		return false
	} else {
		p := new(postion)
		p.value = value
		p.key = key
		list := make([]postion, 1)
		list[0] = *p
		this.List = append(list, this.List...)
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
		return this.List[0].value
	}
}

func (this *MyCircularQueue) Rear() int {
	if !this.IsEmpty() {
		return this.List[len(this.List)-1].value
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

func nextGreaterElements(nums []int) []int {
	c := Constructor(len(nums))
	result := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		if c.IsEmpty() {
			c.EnQueue(nums[i], i)
		} else {
			for c.Front() < nums[i] && len(c.List) != 0 {
				result[c.List[0].key] = nums[i]
				c.DeQueue()
			}
			c.EnQueue(nums[i], i)
		}
	}
	for i := 0; i < len(c.List); i++ {

		for _, num := range nums {
			fmt.Println(nums[i], num)
			if num > c.List[i].value {
				result[c.List[i].key] = num
				goto here
			}
			result[c.List[i].key] = -1
		}
	here:
	}
	return result
}
