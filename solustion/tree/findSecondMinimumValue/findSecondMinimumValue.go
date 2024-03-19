package findSecondMinimumValue

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Min struct {
	min       int
	secondMin int
}

func findSecondMinimumValue(root *TreeNode) int {
	m := new(Min)
	fmt.Println(m.min, m.secondMin)
	m.findMinNode(root)
	if m.secondMin == 0 {
		return -1
	}
	return m.secondMin
}

func (m *Min) findMinNode(root *TreeNode) {
	if root != nil {
		if m.min == 0 {
			m.min = root.Val
		} else {
			if m.min > root.Val {
				m.secondMin = m.min
				m.min = root.Val
			} else if (root.Val < m.secondMin || m.secondMin == 0) && root.Val != m.min {
				m.secondMin = root.Val
			}
		}
		if root.Left != nil {
			m.findMinNode(root.Left)
		}
		if root.Right != nil {
			m.findMinNode(root.Right)
		}
	}
	return
}
