package lowestCommonAncestor

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p == root || q == root {
		return root
	}
	bigger, smaller := new(TreeNode), new(TreeNode)
	if p.Val > q.Val {
		bigger = p
		smaller = q
	} else {
		smaller = p
		bigger = q
	}
	return getAncestor(root, smaller, bigger)
}

func getAncestor(root, smaller, bigger *TreeNode) *TreeNode {

	for root != nil {
		if root.Val >= smaller.Val && root.Val <= bigger.Val {
			return root
		}
		if root.Val >= bigger.Val {
			return getAncestor(root.Left, smaller, bigger)
		}
		if root.Val <= bigger.Val {
			return getAncestor(root.Right, smaller, bigger)
		}
	}
	return root
}
