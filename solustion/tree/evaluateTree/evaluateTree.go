package evaluateTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func evaluateTree(root *TreeNode) bool {
	return evaluate(root)
}

func evaluate(root *TreeNode) bool {
	if root.Val <= 1 {
		return iToB(root.Val)
	} else {
		left, right := root.Left, root.Right
		if left.Val >= 1 {
			left.Val = bToI(evaluate(left))
		}
		if right.Val >= 1 {
			right.Val = bToI(evaluate(right))
		}
		if root.Val == 2 {
			return iToB(right.Val) || iToB(left.Val)
		} else {
			return iToB(right.Val) && iToB(left.Val)
		}
	}
}

func bToI(b bool) int {
	if b {
		return 1
	} else {
		return 0
	}
}

func iToB(num int) bool {
	if num == 0 {
		return false
	} else {
		return true
	}
}
