package simplifyPath

type stacks struct {
	list   []string
	length int
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

func SimplifyPath(path string) string {
	strList, last, left := make([]string, 0), "", false

	for i := 0; i < len(path); i++ {
		if path[i] == 47 {
			if !left {
				left = !left
			} else {
				if last != "" {
					strList = append(strList, last)
				}
				last = ""
			}
		} else {
			last += string(path[i])
		}
	}
	if last != "" {
		strList = append(strList, last)
	}
	s := Init()
	for i := 0; i < len(strList); i++ {
		switch strList[i] {
		case "..":
			s.pop()
		case ".":
			continue
		default:
			s.push(strList[i])
		}
	}
	if s.isZero() {
		return "/"
	}
	result := ""
	for !s.isZero() {
		result += "/" + s.pop()
	}
	return result
}
