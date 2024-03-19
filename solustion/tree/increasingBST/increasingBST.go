package increasingBST

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type InorderTree struct {
	list []int
}

func InorderInit() InorderTree {
	I := new(InorderTree)
	I.list = make([]int, 0)
	return *I
}

func inorderTraversal(root *TreeNode) []int {
	I := InorderInit()
	return I.list
}

func (I *InorderTree) Inorder(root *TreeNode) {
	if root.Left != nil {
		I.Inorder(root.Left)
		I.list = append(I.list, root.Val)
		I.Inorder(root.Right)
	}
	return
}

func increasingBST(root *TreeNode) *TreeNode {
	list := inorderTraversal(root)
	fmt.Println(list)
	tree := new(TreeNode)
	result := tree
	for i := 0; i < len(list)-1; i++ {
		node := new(TreeNode)
		node.Val = list[i]
		tree.Right = node
		tree = node
	}
	return result.Right
}
