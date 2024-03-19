package minAddToMakeValid

func minAddToMakeValid(s string) int {
	left, right := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == 40 {
			left++
		} else {
			if left > 0 {
				left--
			} else {
				right++
			}
		}
	}
	return left + right
}

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
