package diameterOfBinaryTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	l, r := new(count), new(count)
	l.maxDepth(root.Left)
	r.maxDepth(root.Right)

	return max(l.max, r.max)
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

type count struct {
	max int
	now int
}

func (c *count) maxDepth(root *TreeNode) {
	if root != nil {
		c.now++
		if c.now > c.max {
			c.max = c.now
		}
		c.maxDepth(root.Left)
		c.maxDepth(root.Right)
		c.now--
	}
	return
}
