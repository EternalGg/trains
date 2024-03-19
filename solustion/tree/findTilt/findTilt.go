package findTilt

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type postOrder struct {
	m     map[*TreeNode]int
	count int
	max   int
}

func findTilt(root *TreeNode) int {
	p := new(postOrder)
	p.m = make(map[*TreeNode]int)
	p.max = 0
	p.count = 0
	p.PostOrderSearch(root)
	return p.max
}

func (p *postOrder) PostOrderSearch(root *TreeNode) {
	for root != nil {
		p.PostOrderSearch(root.Left)
		p.PostOrderSearch(root.Right)
		fmt.Println(root.Val)
		p.m[root] = root.Val
		p.m[root] += p.m[root.Right] + p.m[root.Left]
		p.max += int(math.Abs(float64(p.m[root.Right] - p.m[root.Left])))
	}
}
