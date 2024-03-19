package solustion

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func TreeSearch(root *TreeNode, target int) *TreeNode {
	if root != nil {
		if root.Val == target {
			return root
		}
		TreeSearch(root.Left, target)
		TreeSearch(root.Right, target)
	}
	return nil
}
