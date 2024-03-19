package calculate

import (
	"fmt"
	"strconv"
)

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

func Calculate(s string) int {
	st := Init()
	last := ""
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
		case '(':
			st.push("(")
		case ')':
			st.push(last)
			l := make([]string, 0)
			for st.head() != "(" {
				now := st.pop()
				l = append(l, now)
			}
			st.pop()
			result := 0
			for j := len(l) - 1; j >= 0; j-- {
				switch l[j] {
				case "+":
					continue
				case "-":
					num, _ := strconv.ParseInt(l[j-1], 10, 64)
					result -= int(num)
					j--
				default:
					num, _ := strconv.ParseInt(l[j], 10, 64)
					result += int(num)
				}
			}
			st.push(strconv.Itoa(result))
			last = ""
		default:
			last += string(s[i])
		}
	}
	if last != "" {
		st.push(last)
	}
	result := 0
	l := make([]string, 0)
	for !st.isZero() {
		now := st.pop()
		l = append(l, now)
	}
	for j := len(l) - 1; j >= 0; j-- {
		switch l[j] {
		case "+":
			num, _ := strconv.ParseInt(l[j-1], 10, 64)
			result += int(num)
			j--
		case "-":
			num, _ := strconv.ParseInt(l[j-1], 10, 64)
			result -= int(num)
			j--
		default:
			num, _ := strconv.ParseInt(l[j], 10, 64)
			result += int(num)
		}
	}
	fmt.Println(result)
	return result
}
