package stack

type bStack struct {
	sList  []byte
	last   byte
	length int
}

func (b *bStack) InitBStack() {
	b.sList = make([]byte, 0)
	b.length = 0
}

func (b *bStack) push(value byte) {
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

func (b *bStack) getLast() byte {
	if b.isZero() {
		return 0
	} else {
		return b.sList[b.length-1]
	}
}
