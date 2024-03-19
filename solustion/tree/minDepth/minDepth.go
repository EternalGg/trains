package minDepth

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Count struct {
	c   int
	min int
}

func minDepth(root *TreeNode) int {
	c := new(Count)
	c.c = 0
	c.countAll(root)
	return c.min
}

func (c *Count) countAll(root *TreeNode) {
	c.c++
	if root != nil {
		if root.Right == nil && root.Left == nil {
			if c.min == 0 {
				c.min = c.c
			} else if c.min < c.c {
				c.min = c.c
			}
		}
		c.countAll(root.Left)
		c.countAll(root.Right)
	}
	c.c--
	return
}
