package bstFromPreorder

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstFromPreorder(preorder []int) *TreeNode {
	tree := new(TreeNode)
	tree.Val = preorder[0]
	for i := 1; i < len(preorder); i++ {
		InsertTree(tree, preorder[i])
	}
	return tree
}

func InsertTree(root *TreeNode, value int) {
	if value > root.Val {
		if root.Right == nil {
			newTree := new(TreeNode)
			newTree.Val = value
			root.Right = newTree
		} else {
			InsertTree(root.Right, value)
		}
	} else {
		if root.Left == nil {
			newTree := new(TreeNode)
			newTree.Val = value
			root.Left = newTree
		} else {
			InsertTree(root.Left, value)
		}
	}
}
