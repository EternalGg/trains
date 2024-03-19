package countNodes

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type PreTree struct {
	list int
}

func PreTreeInit() PreTree {
	p := new(PreTree)
	p.list = 0
	return *p
}

func preorderTraversal(root *TreeNode) int {
	p := PreTreeInit()
	p.preorder(root)
	return p.list
}

func (t *PreTree) preorder(root *TreeNode) {
	if root != nil {
		t.list++
		t.preorder(root.Left)
		t.preorder(root.Right)
	}
	return
}

func countNodes(root *TreeNode) int {
	return preorderTraversal(root)
}
