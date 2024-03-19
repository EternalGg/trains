package MinStack

type tn struct {
	Val  int
	next *tn
}

type MinStack struct {
	Nullnode bool
	node     *tn
	list     []int
	length   int
}

func Constructor() MinStack {
	m := new(MinStack)
	m.node = new(tn)
	m.Nullnode = true
	m.list = make([]int, 0)
	m.length = 0
	return *m
}

func (this *MinStack) Push(val int) {
	this.list = append(this.list, val)
	this.length++
	n := this.node
	if this.Nullnode == true {
		this.node.Val = val
		this.Nullnode = false
		return
	}
	if val <= n.Val {
		newnode := tn{Val: val, next: n}
		this.node = &newnode
		return
	}
	for {
		if n.next == nil {
			newnode := tn{Val: val, next: n.next}
			n.next = &newnode
			return
		} else {
			if val >= n.Val && val < n.next.Val {
				newnode := tn{Val: val, next: n.next}
				n.next = &newnode
				return
			} else {
				n = n.next
			}
		}
	}
}

func (this *MinStack) Pop() {
	if this.length == 0 {
		return
	}
	value := this.list[this.length-1]
	this.list = this.list[:this.length-1]
	this.length--
	if this.node.Val == value {
		if this.node.next == nil {
			this.node.Val = 0
			this.Nullnode = true
			return
		} else {
			this.node = this.node.next
			return
		}
	}
	n := this.node
	for {
		if n.next.Val == value {
			if n.next.next == nil {
				n.next = nil
				return
			} else {
				temple := n.next.next
				n.next = temple
				return
			}
		} else {
			n = n.next
		}
	}
}

func (this *MinStack) Top() int {
	return this.list[this.length-1]
}

func (this *MinStack) GetMin() int {
	return this.node.Val
}
