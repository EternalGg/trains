package isUnivalTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isUnivalTree(root *TreeNode) bool {
	num := root.Val
	list := preorderTraversal(root)
	for _, i2 := range list {
		if i2 != num {
			return false
		}
	}
	return true
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
		t.list = append(t.list, root.Val)
		t.preorder(root.Left)
		t.preorder(root.Right)
	}
	return
}
