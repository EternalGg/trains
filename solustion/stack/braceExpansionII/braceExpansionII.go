package braceExpansionII

type stacks struct {
	list   []string
	length int
	strMap map[string]bool
}

func (s *stacks) head() string {
	return s.list[s.length-1]
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

func BraceExpansionII(expression string) []string {
	st, last := Init(), ""
	for i := 0; i < len(expression); i++ {
		switch expression[i] {
		case '{':
			if last != "" {
				st.push(last)
				last = ""
			}
			st.push("{")
		case '}':
			if last != "" {
				st.push(last)
				last = ""
			}
			tempo := make([]string, 0)
			for !st.isZero() {
				switch st.head() {
				case "{":
					st.pop()

					for next := st.pop(); next != ""; {
						for j := 0; j < len(tempo); j++ {
							if tempo[j] != "," {
								tempo[j] = next + tempo[j]
							}
						}
					}
					for j := 0; j < len(tempo); j++ {
						st.push(tempo[j])
					}
					goto here
				default:
					tempo = append(tempo, st.pop())
				}
			}
		case ',':
			st.push(last)
			st.push(",")
			last = ""
		default:
			last += string(expression[i])
		}
	here:
	}
	return st.list
}
