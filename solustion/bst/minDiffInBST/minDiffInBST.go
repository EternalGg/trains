package minDiffInBST

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
		t.list = append(t.list, root.Val)
		t.preorder(root.Left)
		t.preorder(root.Right)
	}
	return
}

func minDiffInBST(root *TreeNode) int {
	list := preorderTraversal(root)
	min := list[1] - list[0]
	for i := 0; i < len(list)-1; i++ {
		if list[i+1]-list[i] < min {
			min = list[i+1] - list[i]
		}
	}
	return min
}
