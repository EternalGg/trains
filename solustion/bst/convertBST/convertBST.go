package convertBST

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type PreTree struct {
	list []int
}

func PreTreeInit() PreTree {
	p := new(PreTree)
	p.list = make([]int, 0)
	return *p
}

func preorderTraversal(root *TreeNode) []int {
	p := PreTreeInit()
	p.preorder(root)
	return p.list
}

func (t *PreTree) preorder(root *TreeNode) {
	if root != nil {
		t.preorder(root.Left)
		t.list = append(t.list, root.Val)
		t.preorder(root.Right)
	}
	return
}

func preCount(root *TreeNode, M map[int]int) {
	if root != nil {
		root.Val = M[root.Val]
		preCount(root.Left, M)
		preCount(root.Right, M)
	}
	return
}

func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	list := preorderTraversal(root)
	cList, cMap := make([]int, len(list)), make(map[int]int)
	cList[len(list)-1] = list[len(list)-1]
	cMap[list[len(list)-1]] = list[len(list)-1]
	for i := len(list) - 2; i >= 0; i-- {
		cList[i] = cList[i+1] + list[i]
		cMap[list[i]] = cList[i]
	}
	fmt.Println(cList)
	preCount(root, cMap)
	return root
}
