package sumEvenGrandparent

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Count struct {
	Value int
}

func sumEvenGrandparent(root *TreeNode) int {
	c := new(Count)
	c.Value = 0
	c.countAllGrandSon(root)
	return c.Value
}

func (c *Count) countAllGrandSon(root *TreeNode) {
	if root != nil {
		if root.Val%2 == 0 {
			if root.Left != nil {
				if root.Left.Left != nil {
					c.Value += root.Left.Left.Val
				}
				if root.Left.Right != nil {
					c.Value += root.Left.Right.Val
				}
			}
			if root.Right != nil {
				if root.Right.Left != nil {
					c.Value += root.Right.Left.Val
				}
				if root.Right.Right != nil {
					c.Value += root.Right.Right.Val
				}
			}
		}
		c.countAllGrandSon(root.Left)
		c.countAllGrandSon(root.Right)
	}
	return
}
