package tree2str

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func tree2str(root *TreeNode) string {
	now := ""
	if root.Left != nil {
		now += "(" + tree2str(root.Left) + ")"
	} else if root.Right != nil {
		now += "(" + ")"
	}
	if root.Right != nil {
		now += "(" + tree2str(root.Right) + ")"
	}
	return strconv.Itoa(root.Val) + now
}

//func (w *word) postSpell(root *TreeNode) {
//	if {
//
//	}
//}
