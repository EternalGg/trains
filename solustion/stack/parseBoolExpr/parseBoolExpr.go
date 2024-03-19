package parseBoolExpr

type bStack struct {
	sList  []byte
	length int
}

func InitBStack() bStack {
	b := new(bStack)
	b.sList = make([]byte, 0)
	b.length = 0
	return *b
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

func (b *bStack) pop() byte {
	if b.isZero() {
		return 0
	}
	result := b.sList[b.length-1]
	b.length--
	b.sList = b.sList[:b.length]
	return result
}

func ParseBoolExpr(expression string) bool {
	bs := InitBStack()
	for i := 0; i < len(expression); i++ {
		switch expression[i] {
		case ',':
			continue
		case ')':
			tempolist := make([]byte, 0)
			for !bs.isZero() {
				value := bs.pop()
				switch value {
				case '!':
					if tempolist[0] == 't' {
						bs.push('f')
					} else {
						bs.push('t')
					}
					goto b
				case '&':
					for j := 0; j < len(tempolist); j++ {
						if tempolist[j] == 'f' {
							bs.push('f')
							goto b
						}
					}
					bs.push('t')
					goto b
				case '|':
					for j := 0; j < len(tempolist); j++ {
						if tempolist[j] == 't' {
							bs.push('t')
							goto b
						}
					}
					bs.push('f')
					goto b
				case '(':
					continue
				default:
					tempolist = append(tempolist, value)
				}
			}
		default:
			bs.push(expression[i])
		}
	b:
	}
	if bs.pop() == 't' {
		return true
	} else {
		return false
	}
}
