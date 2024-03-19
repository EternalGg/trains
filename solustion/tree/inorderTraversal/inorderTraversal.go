package inorderTraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type InorderTree struct {
	list []int
}

func InorderInit() InorderTree {
	I := new(InorderTree)
	I.list = make([]int, 0)
	return *I
}

func inorderTraversal(root *TreeNode) []int {
	I := InorderInit()
	return I.list
}

func (I *InorderTree) Inorder(root *TreeNode) {
	if root.Left != nil {
		I.Inorder(root.Left)
		I.list = append(I.list, root.Val)
		I.Inorder(root.Right)
	}
	return
}
