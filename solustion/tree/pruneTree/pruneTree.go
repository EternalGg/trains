package pruneTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pruneTree(root *TreeNode) *TreeNode {
	if root.Left != nil {
		root.Left = pruneTree(root.Left)
	}
	if root.Right != nil {
		root.Right = pruneTree(root.Right)
	}
	if root.Left != nil {
		if root.Left.Val == 0 && root.Left.Left == nil && root.Left.Right == nil {
			root.Left = nil
		}
	}
	if root.Right != nil {
		if root.Right.Val == 0 && root.Right.Left == nil && root.Right.Right == nil {
			root.Right = nil
		}
	}
	if root.Left == nil && root.Right == nil && root.Val == 0 {
		return nil
	}
	return root
}
