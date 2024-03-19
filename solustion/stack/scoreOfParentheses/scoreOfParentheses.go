package scoreOfParentheses

type bStack struct {
	sList  []int
	length int
}

func InitBStack() bStack {
	b := new(bStack)
	b.sList = make([]int, 0)
	b.length = 0
	return *b
}

func (b *bStack) push(value int) {
	b.sList = append(b.sList, value)
	b.length++
}

func (b *bStack) isZero() bool {
	if b.length == 0 {
		return true
	} else {
		return false
	}
}

func (b *bStack) getLast() int {
	if b.isZero() {
		return 0
	} else {
		return b.sList[b.length-1]
	}
}

func (b *bStack) pop() int {
	if b.isZero() {
		return 0
	}
	result := b.sList[b.length-1]
	b.length--
	b.sList = b.sList[:b.length]
	return result
}

func ScoreOfParentheses(s string) int {
	bs := InitBStack()
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case ')':
			count := 0
			for !bs.isZero() {
				now := bs.pop()
				switch now {
				case -1:
					if count == 0 {
						bs.push(1)
					} else {
						bs.push(count * 2)
					}
					goto here
				default:
					count += now
				}
			}
		default:
			bs.push(-1)
		}
	here:
	}
	if bs.length != 1 {
		count := 0
		for i := 0; i < len(bs.sList); i++ {
			count += bs.sList[i]
		}
		return count
	} else {
		return bs.getLast()
	}
}
