package countStudents

type stacks struct {
	list   []int
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
	s.length++
}

func (s *stacks) pop() int {
	if s.isZero() {
		return -1
	}
	result := s.list[s.length-1]
	s.list = s.list[:s.length-1]
	s.length--
	return result
}

func Init() stacks {
	s := new(stacks)
	s.list = make([]int, 0)
	s.length = 0
	return *s
}

//
//func countStudents(students []int, sandwiches []int) int {
//	sStack := Init()
//	sStack.list = students
//	sStack.length = sStack.length
//	for i := 0; i < len(sandwiches); i++ {
//		for students[0] != sandwiches[i] {
//			time := 0
//			for time == sStack.length {
//				sStack.push(sStack.list[0])
//				sStack.list = sStack.list[1:]
//				time++
//			}
//		}
//
//			sStack.list = sStack.list[1:]
//			sStack.length--
//		}
//	}
//}
