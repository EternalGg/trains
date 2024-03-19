package sumOfLeftLeaves

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Counter struct {
	count int
}

func sumOfLeftLeaves(root *TreeNode) int {
	c := new(Counter)
	c.addLeft(root)
	return c.count
}

func (c *Counter) addLeft(root *TreeNode) {
	if root != nil {
		if root.Left != nil {
			if root.Left.Left == nil && root.Left.Right == nil {
				c.count += root.Left.Val
			}
		}
		c.addLeft(root.Left)
		c.addLeft(root.Right)
	}
	return
}
