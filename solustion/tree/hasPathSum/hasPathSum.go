package hasPathSum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type PreTree struct {
	target int
}

func PreTreeInit() PreTree {
	p := new(PreTree)
	p.target = 0
	return *p
}

func (t *PreTree) preorder(root *TreeNode) bool {
	if root != nil {
		t.target -= root.Val
		if t.target == 0 {
			return true
		}
		if t.preorder(root.Left) {
			return true
		}
		if t.preorder(root.Right) {
			return true
		}
		t.target += root.Val
	}
	return false
}

func HasPathSum(root *TreeNode, targetSum int) bool {
	p := PreTreeInit()
	p.target = targetSum
	result := p.preorder(root)

	return result
}
