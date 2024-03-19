package inorderSuccessor

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	if root != nil {
		if root.Left == nil && root.Val == p.Val {
			return root
		}
	} else if root.Left != nil {
		if root.Left.Val == p.Val {
			return root
		} else {
			return inorderSuccessor(root.Left, p)
		}
	} else if root.Right != nil {
		return inorderSuccessor(root.Right, p)
	}
	return nil
}
