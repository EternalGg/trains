package findTarget

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type PreTree struct {
	iMap map[int]bool
	list []int
}

func PreTreeInit() PreTree {
	p := new(PreTree)
	p.list = make([]int, 0)
	p.iMap = make(map[int]bool)
	return *p
}

func preorderTraversal(root *TreeNode) ([]int, map[int]bool) {
	p := PreTreeInit()
	p.preorder(root)
	return p.list, p.iMap
}

func (t *PreTree) preorder(root *TreeNode) {
	if root != nil {
		t.list = append(t.list, root.Val)
		t.iMap[root.Val] = true
		t.preorder(root.Left)
		t.preorder(root.Right)
	}
	return
}

func findTarget(root *TreeNode, k int) bool {
	list, iMAp := preorderTraversal(root)
	for i := 0; i < len(list); i++ {
		if iMAp[k-list[i]] && k-list[i] != list[i] {
			return true
		}
	}
	return false
}
