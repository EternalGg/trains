package checkSubTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func checkSubTree(t1 *TreeNode, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	if t1 != nil {
		if t1.Val == t2.Val {
			if preorderTwoTree(t1, t2) {
				return true
			}
		}
	}
	return checkSubTree(t1.Left, t2) || checkSubTree(t1.Right, t2)
}

func preorderTwoTree(root1, root2 *TreeNode) bool {

	if !(root1.Left == nil && root2.Left == nil) {
		return preorderTwoTree(root1.Left, root2.Left)
	}
	if !(root1.Right == nil && root2.Right == nil) {
		return preorderTwoTree(root1.Right, root2.Right)
	}

	return true
}
