package isBalanced

import (
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type treeDigger struct {
	tMap    map[*TreeNode]int
	balance bool
}

func isBalanced(root *TreeNode) bool {
	td := new(treeDigger)
	td.tMap = make(map[*TreeNode]int)
	td.balance = true
	td.Digger(root)

	return td.balance
}

func (t *treeDigger) Digger(root *TreeNode) {
	if root != nil {
		t.Digger(root.Left)
		t.Digger(root.Right)
	}
	if root == nil {
		return
	} else {
		t.tMap[root] = 1
		switch {
		case root.Left == nil && root.Right == nil:
		case root.Left == nil:
			if t.tMap[root.Right] > 1 {
				t.balance = false
			} else {
				t.tMap[root] += t.tMap[root.Right]
			}
		case root.Right == nil:
			if t.tMap[root.Left] > 1 {
				t.balance = false
			} else {
				t.tMap[root] += t.tMap[root.Left]
			}
		default:
			if math.Abs(float64(t.tMap[root.Left])-float64(t.tMap[root.Right])) > 1 {
				t.balance = false
			} else {
				t.tMap[root] += max(t.tMap[root.Left], t.tMap[root.Right])
			}
		}
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
