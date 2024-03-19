package isPalindrome

type ListNode struct {
	Val  int
	Next *ListNode
}

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

func isPalindrome(head *ListNode) bool {
	s := Init()
	slow, fast := head, head
	for fast.Next != nil {
		if fast.Next == nil {
			break
		} else {
			s.push(slow.Val)
			fast = fast.Next
		}
		if fast.Next == nil {
			break
		} else {
			fast = fast.Next
		}
		slow = slow.Next
	}
	for slow != nil {
		if slow.Val != s.head() {
			return false
		}
		slow = slow.Next
		s.pop()
	}
	return true
}
