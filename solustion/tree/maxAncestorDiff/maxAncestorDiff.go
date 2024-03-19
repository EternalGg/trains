package maxAncestorDiff

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type save struct {
	list []int
	max  int
	min  int
	Max  float64
}

func maxAncestorDiff(root *TreeNode) int {
	s := new(save)
	s.list = make([]int, 0)
	s.min = root.Val
	return int(s.Max)
}

func (s *save) dfs(root *TreeNode) {
	if root != nil {

		s.list = append(s.list, root.Val)
		if root.Val > s.max {
			s.max = root.Val
		}
		if root.Val < s.max {
			s.min = root.Val
		}
		if math.Abs(float64(s.min)-float64(root.Val)) > s.Max {
			s.Max = math.Abs(float64(s.min) - float64(root.Val))
		}
		if math.Abs(float64(s.max)-float64(root.Val)) > s.Max {
			s.Max = math.Abs(float64(s.max) - float64(root.Val))
		}
		s.dfs(root.Left)
		s.dfs(root.Right)
		if root.Val == s.min || root.Val == s.max {
			if root.Val == s.max {
				s.list = s.list[:len(s.list)]
				s.findMax()
			} else {
				s.list = s.list[:len(s.list)]
				s.findMin()
			}
		} else {
			s.list = s.list[:len(s.list)]
		}
	}
}

func (s *save) findMin() {
	m := s.list[0]
	for i := 0; i < len(s.list); i++ {
		if s.list[i] < m {
			m = s.list[i]
		}
	}
	s.min = m
}

func (s *save) findMax() {
	m := s.list[0]
	for i := 0; i < len(s.list); i++ {
		if s.list[i] > m {
			m = s.list[i]
		}
	}
	s.Max = float64(m)
}
