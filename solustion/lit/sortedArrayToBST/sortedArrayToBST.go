package sortedArrayToBST

import "fmt"

func sortedArrayToBST(nums []int) *TreeNode {
	half := len(nums) / 2
	tree := new(TreeNode)
	tree.Val = nums[half]

	for i := 0; i >= half; i-- {
		insertIntoBST(tree, nums[i])
		fmt.Println(i)
		insertIntoBST(tree, nums[len(nums)-i-1])
		fmt.Println(len(nums) - i - 1)
	}
	return tree
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root.Val == val {
		return root
	}
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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
