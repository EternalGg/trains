package averageOfSubtree

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type find struct {
	max   int
	now   int
	count int
}

func averageOfSubtree(root *TreeNode) int {
	l := new(find)
	r := new(find)
	l.findMax(root.Left)
	r.findMax(root.Right)
	fmt.Println(r.max, l.max)
	if (l.now+r.now+root.Val)/(l.count+r.count+1) == root.Val {
		return l.max + r.max + 1
	}
	return l.max + r.max
}

func (f *find) findMax(root *TreeNode) {
	if root.Left != nil {
		f.findMax(root.Left)
		f.findMax(root.Right)
		f.now += root.Val
		if f.now/f.count == root.Val {
			f.max++
		}
	}
	f.count++
}
