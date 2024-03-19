package validatestacksSequences

type stacks struct {
	list   []int
	eMap   map[int]bool
	length int
}

func (s *stacks) head() int {
	return s.list[s.length-1]
}

func (s *stacks) isZero() bool {
	if len(s.list) == 0 {
		return true
	} else {
		return false
	}
}

func (s *stacks) push(val int) {
	s.list = append(s.list, val)
	s.eMap[val] = true
	s.length++
}

func (s *stacks) pop() int {
	result := s.list[s.length-1]
	s.list = s.list[:s.length-1]
	s.length--
	return result
}

func Init() stacks {
	s := new(stacks)
	s.list = make([]int, 0)
	s.length = 0
	s.eMap = make(map[int]bool)
	return *s
}

func validatestacksSequences(pushed []int, popped []int) bool {
	s := new(stacks)
	key := 0
	for i := 0; i < len(popped); i++ {
		if s.isZero() && key != len(pushed) {
			s.push(pushed[key])
			key++
		}
		if s.head() != popped[i] {
			if s.eMap[popped[i]] == false {
				for popped[i] != s.head() {
					s.push(pushed[key])
					key++
				}
				s.pop()
			} else {
				return false
			}
		} else {
			s.pop()
		}
	}

	return true
}
