package addToArrayForm

import "strconv"

type Queue struct {
	List   []int
	length int
}

func initQueue() Queue {
	queue := new(Queue)
	queue.List = make([]int, 0)
	queue.length = 0
	return *queue
}

func (q *Queue) push(num int) {
	q.List = append(q.List, num)
	q.length++

}

func (q *Queue) pop() int {
	if q.length == 0 {
		return 0
	}
	result := q.List[q.length-1]
	q.List = q.List[:q.length-1]
	q.length--
	return result
}

func AddToArrayForm(num []int, k int) []int {
	n2 := numToArray(k)
	n1Queue := initQueue()
	for _, value := range num {
		n1Queue.push(value)
	}
	n2Queue := initQueue()
	for _, value := range n2 {
		n2Queue.push(value)
	}
	result, up := make([]int, 0), false

	for n1Queue.length != 0 || n2Queue.length != 0 {
		a1, a2 := n1Queue.pop(), n2Queue.pop()
		plus := a1 + a2
		if up {
			plus++
		}
		add := plus
		if plus >= 10 {
			add = plus % 10
			up = true
		} else {
			up = false
		}
		result = append(result, add)
	}

	if up {
		result = append(result, 1)
	}
	result = reverseList(result)
	return result
}

func reverseList(numList []int) []int {
	result := make([]int, len(numList))
	for i := 0; i <= len(numList)-1; i++ {
		result[len(numList)-i-1] = numList[i]
	}
	return result
}

func numToArray(num int) []int {
	str := strconv.Itoa(num)
	result := make([]int, 0)
	for _, value := range str {
		now := value - 48
		result = append(result, int(now))
	}
	return result
}
