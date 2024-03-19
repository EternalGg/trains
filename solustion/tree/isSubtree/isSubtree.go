package isSubtree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	return dfs1(root, subRoot)
}

func dfs1(roots *TreeNode, subRoot *TreeNode) bool {
	if roots == nil {
		return false
	} else if dfs(roots, subRoot) {
		return true
	} else {
		return dfs1(roots.Left, subRoot) || dfs1(roots.Right, subRoot)
	}
}

func dfs(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	} else if root == nil || subRoot == nil {
		return false
	} else if root.Val != subRoot.Val {
		return false
	} else {
		return dfs(root.Left, subRoot.Left) && dfs(root.Right, subRoot.Right)
	}
}
