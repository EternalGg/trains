package removeLeafNodes

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	if root.Left != nil {
		if root.Left != nil && root.Left.Val == target && root.Left.Left.Val == target && root.Left.Right.Val == target {
			fmt.Println("find_left")
			root.Left = nil
		} else {
			fmt.Println(root.Val)
			removeLeafNodes(root.Left, target)
		}
	} else if root.Right != nil {
		if root.Right != nil && root.Right.Val == target && root.Right.Left.Val == target && root.Right.Right.Val == target {
			fmt.Println("find_right")
			root.Right = nil
		} else {
			fmt.Println(root.Val)
			root.Right = nil
		}
	}
	return root
}
