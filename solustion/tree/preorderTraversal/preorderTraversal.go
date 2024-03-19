package preorderTraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type PreTree struct {
	list []int
}

func PreTreeInit() PreTree {
	p := new(PreTree)
	p.list = make([]int, 0)
	return *p
}

func preorderTraversal(root *TreeNode) []int {
	p := PreTreeInit()
	p.preorder(root)
	return p.list
}

func (t *PreTree) preorder(root *TreeNode) {
	if root != nil {
		t.preorder(root.Left)
		t.list = append(t.list, root.Val)
		t.preorder(root.Right)
	}
	return
}
