package reverseStr

import "fmt"

type Queue struct {
	list []uint8
}

func (q *Queue) Push(num uint8) {
	q.list = append(q.list, num)
}
func (q *Queue) Pop() string {
	result := string(q.list)
	q.list = make([]uint8, 0)
	return result
}

func reverseStr(s string, k int) string {
	isVerse, result, key, q := true, "", 1, new(Queue)
	q.list = make([]uint8, 0)
	for i := 0; i < len(s); i++ {
		if i == (key*k)-1 {
			fmt.Println(i)
			key++
			str := q.Pop()
			if isVerse {
				str = Reverse(str)
			}
			result += str
			isVerse = !isVerse
		} else {
			fmt.Println(i)
			q.Push(s[i])
		}
	}
	return result
}

func Reverse(s string) string {
	result := ""
	for i := len(s) - 1; i >= 0; i-- {
		result += string(s[i])
	}
	return result
}
