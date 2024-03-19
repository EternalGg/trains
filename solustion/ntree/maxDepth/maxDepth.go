package maxDepth

type Node struct {
	Val      int
	Children []*Node
}

type depth struct {
	deep int
	max  int
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	d := new(depth)
	d.deep = 1
	d.depthDigger(root)
	return d.max
}

func (d *depth) depthDigger(root *Node) {
	if root != nil {
		for i := 0; i < len(root.Children); i++ {
			d.deep++
			if d.deep > d.max {
				d.max = d.deep
			}
			d.depthDigger(root.Children[i])
		}
	}
	d.deep--
}
