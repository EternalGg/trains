package leafSimilar

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type LeafList struct {
	List []int
}

func LeafListInit() *LeafList {
	l := new(LeafList)
	l.List = make([]int, 0)
	return l
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	r1l, r2l := LeafListInit(), LeafListInit()
	r1l.preorder(root1)
	r2l.preorder(root2)
	if len(r1l.List) != len(r2l.List) {
		return false
	}
	for i, i2 := range r1l.List {
		if r2l.List[i] != i2 {
			return false
		}
	}
	return true
}

func (l *LeafList) preorder(root *TreeNode) {
	if root != nil {
		if root.Left == nil && root.Right == nil {
			l.List = append(l.List, root.Val)
		}
		l.preorder(root.Left)
		l.preorder(root.Right)
	}
	return
}
