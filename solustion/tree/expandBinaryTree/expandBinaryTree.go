package expandBinaryTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func expandBinaryTree(root *TreeNode) *TreeNode {
	addTree(root)
	return root
}

func addTree(root *TreeNode) *TreeNode {
	if root != nil {
		add(root)
		if root.Left != nil {
			addTree(root.Left.Left)
		}
		if root.Right != nil {
			addTree(root.Right.Right)
		}
	}
	return root
}

func add(root *TreeNode) {
	if root.Left != nil {
		newOne := new(TreeNode)
		newOne.Val = -1
		newOne.Left = root.Left
		root.Left = newOne
	}
	if root.Right != nil {
		newOne := new(TreeNode)
		newOne.Val = -1
		newOne.Right = root.Right
		root.Right = newOne
	}
}
