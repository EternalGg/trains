package minimizedStringLength

type queue struct {
	list   []byte
	length uint
}

func (q *queue) push(i byte) {
	q.list = append(q.list, i)
	q.length++
}

func (q *queue) pop() {
	q.length--
	q.list = q.list[:q.length]
}

func (q *queue) pophead() {

	q.list = q.list[1:q.length]
	q.length--
}
func (q *queue) head() byte {
	return q.list[q.length-1]
}

func MinimizedStringLength(s string) int {
	q := queue{make([]byte, 0), 0}

	for i := 0; i < len(s); i++ {
		for q.length != 0 {
			if q.head() == s[i] {
				q.pop()
			} else if q.list[0] == s[i] {
				q.pophead()
			} else {
				break
			}
		}
		q.push(s[i])
	}
	return int(q.length)
}
