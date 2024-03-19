package minLength

type stacks struct {
	list   []string
	length int
	strMap map[string]bool
}

func (s *stacks) head() string {
	if s.isZero() {
		return "z"
	} else {
		return s.list[s.length-1]
	}
}

func (s *stacks) isZero() bool {
	if len(s.list) == 0 {
		return true
	} else {
		return false
	}
}

func (s *stacks) push(val string) {
	s.list = append(s.list, val)
	s.length++
}

func (s *stacks) pop() string {
	if s.isZero() {
		return ""
	}
	result := s.list[s.length-1]
	s.list = s.list[:s.length-1]
	s.length--
	return result
}

func Init() stacks {
	s := new(stacks)
	s.list = make([]string, 0)
	s.length = 0
	return *s
}

func MinLength(s string) int {
	stack := Init()
	for i := 0; i < len(s); i++ {
		if s[i] == 'B' || s[i] == 'D' {
			if (stack.head() == "A" && s[i] == 'B') || (stack.head() == "C" && s[i] == 'D') {
				stack.pop()
			} else {
				stack.push(string(s[i]))
			}
		} else {
			stack.push(string(s[i]))
		}
	}
	return stack.length
}
