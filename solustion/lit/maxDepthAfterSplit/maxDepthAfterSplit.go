package maxDepthAfterSplit

type stacks struct {
	list   []bool
	length int
}

func (s *stacks) isZero() bool {
	if len(s.list) == 0 {
		return true
	} else {
		return false
	}
}

func (s *stacks) push(val bool) {
	s.list = append(s.list, val)
	s.length++
}

func (s *stacks) pop() {
	if s.isZero() {
		return
	}
	s.list = s.list[:s.length-1]
	s.length--
	return
}

func Init() stacks {
	s := new(stacks)
	s.list = make([]bool, 0)
	s.length = 0
	return *s
}

func maxDepthAfterSplit(seq string) []int {
	is := Init()
	result := make([]int, len(seq))
	for i := 0; i < len(seq); i++ {
		if seq[i] == '(' {
			is.push(true)
			result[i] = is.length - 1
		} else {
			result[i] = is.length - 1
			is.pop()
		}
	}
	return result
}
