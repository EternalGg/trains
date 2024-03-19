package pathSumm

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type stacks struct {
	list   []int
	length int
	target int
	result int
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
	sum := 0
	s.list = append(s.list, val)
	for i := len(s.list) - 1; i >= 0; i-- {
		sum += s.list[i]
		if sum == s.target {
			s.result++
		}
	}
	s.length++
}

func (s *stacks) pop() int {
	if s.isZero() {
		return 0
	}
	result := s.list[s.length-1]
	s.list = s.list[:s.length-1]
	s.length--
	return result
}

func Init(target int) stacks {
	s := new(stacks)
	s.list = make([]int, 0)
	s.target = target
	s.length = 0
	return *s
}

func pathSum(root *TreeNode, sum int) int {
	ts := Init(sum)
	ts.Find(root)
	return ts.result
}

func (ts *stacks) Find(root *TreeNode) {
	if root != nil {
		ts.push(root.Val)
		ts.Find(root.Left)
		ts.Find(root.Right)
		ts.pop()
	}
}
