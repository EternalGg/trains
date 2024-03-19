package addOneRow

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type depp struct {
	depth int
}

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		return &TreeNode{val, root, nil}
	}

	d := new(depp)
	d.depth = 1
	d.dfs(root, val, depth)
	return root
}

func (d *depp) dfs(root *TreeNode, val int, depth int) {
	if d.depth < depth-1 {
		d.depth++
		if root.Left != nil {
			d.dfs(root.Left, val, depth)
		}
		if root.Right != nil {
			d.dfs(root.Right, val, depth)
		}
	} else {

		if root.Left != nil {
			l := root.Left
			root.Left = &TreeNode{val, l, nil}
		}
		if root.Right != nil {
			r := root.Right
			root.Right = &TreeNode{val, nil, r}
		}
		d.depth--
		return
	}

}
