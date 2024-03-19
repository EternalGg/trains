package invertTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorder(root *TreeNode) {
	if root != nil {
		middle := root.Left
		root.Left = root.Right
		root.Right = middle
		preorder(root.Left)
		preorder(root.Right)
	}
	return
}

func invertTree(root *TreeNode) *TreeNode {
	preorder(root)
	return root
}
