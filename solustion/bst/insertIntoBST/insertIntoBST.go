package insertIntoBST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		root = new(TreeNode)
		root.Val = val
		return root
	}
	if val < root.Val {
		left := insertIntoBST(root.Left, val)
		root.Left = left
	}
	if val > root.Val {
		right := insertIntoBST(root.Right, val)
		root.Right = right
	}
	return root
}
