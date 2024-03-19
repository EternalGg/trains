package sortedArrayToBST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		root = new(TreeNode)
		root.Val = val
		return root
	}
	if val < root.Val {
		left := insertIntoBST(root.Left, val)
		root.Left = left
	}
	if val > root.Val {
		right := insertIntoBST(root.Right, val)
		root.Right = right
	}
	return root
}

func sortedArrayToBST(nums []int) *TreeNode {
	tree := new(TreeNode)
	tree.Val = nums[0]
	for i := len(nums) / 2; i < len(nums); i++ {
		insertIntoBST(tree, nums[i])
	}
	for i := len(nums) / 2; i >= 0; i-- {
		insertIntoBST(tree, nums[i])
	}
	return tree
}
