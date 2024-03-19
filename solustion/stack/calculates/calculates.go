package calculates

import "strconv"

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

func (s *stacks) headPop() string {
	if s.isZero() {
		return ""
	}
	result := s.list[0]
	s.list = s.list[1:]
	s.length--
	return result
}

func Init() stacks {
	s := new(stacks)
	s.list = make([]string, 0)
	s.length = 0
	return *s
}

func calculate(s string) int {
	st, last := Init(), ""
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case ' ':
			continue
		case '+':
			st.push(last)
			st.push("+")
			last = ""
		case '-':
			st.push(last)
			st.push("-")
			last = ""
		case '*':
			st.push(last)
			st.push("*")
			last = ""
		case '/':
			st.push(last)
			st.push("/")
			last = ""
		default:

			last += string(s[i])
		}
	}
	if last != "" {
		st.push(last)
	}
	st1 := Init()

	for i := 0; i < len(st.list); i++ {
		now := st.list[i]
		switch now {
		case "*":
			next := st.list[i+1]
			head := st1.pop()
			ni, _ := strconv.ParseInt(next, 10, 64)
			li, _ := strconv.ParseInt(head, 10, 64)
			result := int(ni * li)
			st1.push(strconv.Itoa(result))
			i++
		case "/":
			next := st.list[i+1]
			head := st1.pop()
			ni, _ := strconv.ParseInt(next, 10, 64)
			li, _ := strconv.ParseInt(head, 10, 64)
			result := int(li / ni)
			st1.push(strconv.Itoa(result))
			i++
		default:
			st1.push(now)
		}
	}

	result := 0
	for i := 0; i < len(st1.list); i++ {
		now := st1.list[i]
		switch now {
		case "+":
			next := st1.list[i+1]
			ni, _ := strconv.ParseInt(next, 10, 64)
			result += int(ni)
			i++
		case "-":
			next := st1.list[i+1]
			ni, _ := strconv.ParseInt(next, 10, 64)
			result -= int(ni)
			i++
		default:
			ni, _ := strconv.ParseInt(now, 10, 64)
			result += int(ni)
		}
	}
	return result
}
