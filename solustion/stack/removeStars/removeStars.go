package removeStars

type stacks struct {
	list   []byte
	length int
}

func (s *stacks) head() byte {
	return s.list[s.length-1]
}

func (s *stacks) isZero() bool {
	if len(s.list) == 0 {
		return true
	} else {
		return false
	}
}

func (s *stacks) push(val byte) {
	s.list = append(s.list, val)
	s.length++
}

func (s *stacks) pop() byte {
	if s.isZero() {
		return 0
	}
	result := s.list[s.length-1]
	s.list = s.list[:s.length-1]
	s.length--
	return result
}

func Init() stacks {
	s := new(stacks)
	s.list = make([]byte, 0)
	s.length = 0
	return *s
}

func removeStars(s string) string {
	st := Init()
	b := []byte(s)
	for i := 0; i < len(b); i++ {
		switch b[i] {
		case '*':
			st.pop()
		default:
			st.push(b[i])
		}
	}
	return string(st.list)
}
