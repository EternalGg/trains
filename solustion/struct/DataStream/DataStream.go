package DataStream

type Queue struct {
	List   []int
	Length int
	Max    int
}

func initQueue(max int) *Queue {
	queue := new(Queue)
	queue.List = make([]int, 0)
	queue.Length = 0
	queue.Max = max
	return queue
}

func (q *Queue) push(num int) {
	q.List = append(q.List, num)
	q.Length++
}

func (q *Queue) pop() int {
	if q.Length == 0 {
		return 0
	}
	result := q.List[0]
	q.List = q.List[1:]
	q.Length--
	return result
}

type DataStream struct {
	q     *Queue
	Value int
	m     map[int]int
}

func Constructor(value int, k int) DataStream {
	d := new(DataStream)
	q := initQueue(k)
	d.q = q
	d.Value = value
	d.m = make(map[int]int)
	return *d
}

func (this *DataStream) Consec(num int) bool {
	this.m[num]++
	this.q.push(num)
	if this.q.Max < len(this.q.List) {
		pop := this.q.pop()
		this.m[pop]--
	}
	if this.m[this.Value] == this.q.Max {
		return true
	} else {
		return false
	}
}
