package isSameTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if (p == nil && q != nil) || (q == nil && p != nil) {
		return false
	}
	if p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}
		if !(isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)) {
			return false
		}
	}
	return true
}
