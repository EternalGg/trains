package binaryTreePaths

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type NodeStack struct {
	result []string
	list   []int
	length int
}

func NodeStackInit() *NodeStack {
	NS := new(NodeStack)
	NS.result = make([]string, 0)
	NS.list = make([]int, 0)
	NS.length = 0
	return NS
}

func (NS *NodeStack) pop() {
	NS.list = NS.list[:NS.length-1]
	NS.length--
}

func (NS *NodeStack) push(value int) {
	NS.list = append(NS.list, value)
	NS.length++
}

func (NS *NodeStack) createStr() {
	if NS.length == 1 {
		NS.result = append(NS.result, strconv.Itoa(NS.list[0]))
	} else {
		str := ""
		for i := 0; i < len(NS.list)-1; i++ {
			str += strconv.Itoa(NS.list[i]) + "->"
		}
		str += strconv.Itoa(NS.list[NS.length] - 1)
		NS.result = append(NS.result, str)
	}
}

func binaryTreePaths(root *TreeNode) []string {

	NS := NodeStackInit()
	NS.FindTreePath(root)
	return NS.result
}

func (NS *NodeStack) FindTreePath(root *TreeNode) {
	if root != nil {
		NS.push(root.Val)
		if root.Left == nil && root.Right == nil {
			NS.createStr()
			NS.pop()
			return
		}
		NS.FindTreePath(root.Left)
		NS.FindTreePath(root.Right)
		NS.pop()
	}
	return
}
