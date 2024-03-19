package mergeTrees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {

	return mergeTwoTree(root1, root2)
}

func mergeTwoTree(t1, t2 *TreeNode) *TreeNode {
	if t1 != nil && t2 != nil {
		t1.Val += t2.Val
		if t1.Left == nil && t2.Left != nil {
			t1.Left = t2.Left
		} else {
			mergeTwoTree(t1.Left, t2.Left)
		}
		if t1.Right == nil && t2.Right != nil {
			t1.Right = t2.Right
		} else {
			mergeTwoTree(t1.Right, t2.Right)
		}
	}
	return t1
}
