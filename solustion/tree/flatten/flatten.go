package flatten

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type PreTree struct {
	list []*TreeNode
}

func PreTreeInit() PreTree {
	p := new(PreTree)
	p.list = make([]*TreeNode, 0)
	return *p
}

func preorderTraversal(root *TreeNode) []*TreeNode {
	p := PreTreeInit()
	p.preorder(root)
	return p.list
}

func (t *PreTree) preorder(root *TreeNode) {
	if root != nil {
		t.list = append(t.list, root)
		t.preorder(root.Left)
		t.preorder(root.Right)
	}
	return
}

func flatten(root *TreeNode) {
	list := preorderTraversal(root)
	for i := 0; i < len(list)-1; i++ {
		list[i].Left = nil
		list[i].Right = list[i+1]
	}
	list[len(list)-1].Left = nil
	list[len(list)-1].Right = nil
}
