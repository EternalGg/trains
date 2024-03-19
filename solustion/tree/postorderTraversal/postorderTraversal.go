package postorderTraversal

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

func postorderTraversal(root *TreeNode) []int {
	p := InorderInit()
	p.postorder(root)
	return p.list
}

func (I *InorderTree) postorder(root *TreeNode) {
	if root.Left != nil {
		I.postorder(root.Left)
		I.postorder(root.Right)
		I.list = append(I.list, root.Val)
	}
	return
}
