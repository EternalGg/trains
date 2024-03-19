package isValidBST

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Judge struct {
	b bool
}

func (j *Judge) isValidB(root *TreeNode) {
	if root != nil {
		if root.Left != nil {
			fmt.Println(root.Val, root.Left.Val, "left")
			if root.Left.Val >= root.Val {
				j.b = false
				return
			} else {
				j.isValidB(root.Left)
			}
		}
		if root.Right != nil {
			fmt.Println(root.Val, root.Right.Val, "right")
			if root.Right.Val <= root.Val {
				j.b = false
				return
			} else {
				j.isValidB(root.Right)
			}
		}
	}
}

func isValidBST(root *TreeNode) bool {
	j := new(Judge)
	j.b = true
	j.isValidB(root)
	fmt.Println(j.b)
	if j.b {
		return true
	} else {
		return false
	}
}
