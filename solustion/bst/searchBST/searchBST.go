package searchBST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	return search(root, val)
}

func search(root *TreeNode, val int) *TreeNode {
	if root != nil {
		if root.Val == val {
			return root
		}
		if root.Val < val {
			return search(root.Right, val)
		} else {
			return search(root.Left, val)
		}
	} else {
		return nil
	}
}
