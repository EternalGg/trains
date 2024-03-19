package BSTIterator

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

type BSTIterator struct {
	list []int
	key  int
}

func Constructor(root *TreeNode) BSTIterator {
	b := new(BSTIterator)
	b.list = preorderTraversal(root)
	b.key = 0
	return *b
}

func (this *BSTIterator) Next() int {
	result := this.list[this.key]
	this.key++
	return result
}

func (this *BSTIterator) HasNext() bool {
	if this.key == len(this.list) {
		return false
	} else {
		return true
	}
}
