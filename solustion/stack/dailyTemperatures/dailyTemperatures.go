package dailyTemperatures

type postion struct {
	value int
	key   int
}

type NodeList struct {
	Node postion
	Next *NodeList
}

func Constructor() NodeList {
	m := new(NodeList)
	p := new(postion)
	p.value = -1
	p.key = -1
	m.Node = *p
	m.Next = nil
	return *m
}

func (m *NodeList) Front() int {
	if m.IsEmpty() {
		return -1
	} else {
		return m.Node.value
	}
}

func (m *NodeList) IsEmpty() bool {
	if m.Node.value == -1 {
		return true
	} else {
		return false
	}
}

func DailyTemperatures(temperatures []int) []int {
	c := Constructor()
	head := c
	result := make([]int, len(temperatures))
	for i := 0; i < len(temperatures); i++ {
		for !head.IsEmpty() && head.Front() < temperatures[i] {
			key := head.Node.key
			head = *head.Next
			result[key] = i - key
		}
		z := head
		p := new(postion)
		p.value = temperatures[i]
		p.key = i
		n := new(NodeList)
		n.Node = *p
		n.Next = &z
		head = *n
	}

	return result
}
