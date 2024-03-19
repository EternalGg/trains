package numUniqueEmails

type stacks struct {
	list   []string
	length int
	strMap map[string]bool
}

func (s *stacks) head() string {
	if s.isZero() {
		return ""
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

func numUniqueEmails(emails []string) int {
	result, eMap := 0, make(map[string]bool)
	for _, email := range emails {
		st, isCom, isPlus := Init(), false, false
		for i := 0; i < len(email); i++ {
			if isCom {
				st.push(string(email[i]))
			} else {
				switch email[i] {
				case '.':
					continue
				case '+':
					isPlus = true
				case '@':
					st.push(string(email[i]))
					isCom = true
				default:
					if isPlus {
						continue
					} else {
						st.push(string(email[i]))
					}
				}
			}
		}
		now := ""

		for i := 0; i < len(st.list); i++ {
			now += st.list[i]
		}
		if !eMap[now] {
			result++
		} else {
			eMap[now] = true
		}
	}
	return result
}
