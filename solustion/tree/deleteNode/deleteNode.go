package deleteNode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorder(t *TreeNode, target int) {
	if t != nil {
		if t.Left.Val == target || t.Right.Val == target {
			Delete(t, target)
			return
		}
		// t<目标
		if t.Val < target {
			preorder(t.Right, target)
		} else {
			preorder(t.Left, target)
		}
	}
	return
}

func Delete(t *TreeNode, target int) {
	if t.Left.Val == target {
		if t.Left.Right == nil {
			t.Left = t.Left.Left
		} else if t.Left.Left == nil {
			t.Left = t.Left.Right
		} else {
			t.Left.Right.Left = t.Left.Left
			t.Left = t.Left.Right
		}
	} else {
		if t.Left.Right == nil {
			t.Right = t.Left.Left
		} else if t.Left.Left == nil {
			t.Right = t.Left.Right
		} else {
			t.Right.Right.Left = t.Left.Left
			t.Right = t.Left.Right
		}
	}
}

func DeleteNode(root *TreeNode, key int) *TreeNode {
	preorder(root, key)
	return root
}
