package reformat

import "math"

type Queue struct {
	List   []int32
	length int
}

func initQueue() Queue {
	queue := new(Queue)
	queue.List = make([]int32, 0)
	queue.length = 0
	return *queue
}

func (q *Queue) push(num int32) {
	q.List = append(q.List, num)
	q.length++

}

func (q *Queue) pop() int32 {
	if q.length == 0 {
		return -1
	}
	result := q.List[q.length-1]
	q.List = q.List[:q.length-1]
	q.length--
	return result
}

func Reformat(s string) string {
	nQueue, sQueue := initQueue(), initQueue()

	for _, value := range s {
		if value >= 48 && value <= 57 {
			nQueue.push(value)
		}
		if value >= 97 && value <= 122 {
			sQueue.push(value)
		}
	}
	result := ""
	if !(math.Abs(float64(nQueue.length-sQueue.length)) > 1) {
		switch nQueue.length - sQueue.length {
		// n小 s大
		case -1:
			for nQueue.length != 0 && sQueue.length != 0 {
				result += string(sQueue.pop())
				result += string(nQueue.pop())
			}
			last := sQueue.pop()
			if last != -1 {
				result += string(last)
			}
		default:
			for nQueue.length != 0 && sQueue.length != 0 {
				result += string(nQueue.pop())
				result += string(sQueue.pop())
			}
			last := nQueue.pop()
			if last != -1 {
				result += string(last)
			}
		}
	}
	return result
}
