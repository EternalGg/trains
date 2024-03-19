package preorder

type Node struct {
	Val      int
	Children []*Node
}

type Result struct {
	list []int
}

func preorder(root *Node) []int {
	r := new(Result)
	r.preorderDigger(root)
	return r.list
}

func (r *Result) preorderDigger(root *Node) {
	if root != nil {
		r.list = append(r.list, root.Val)
		for i := 0; i < len(root.Children); i++ {
			r.preorderDigger(root.Children[i])
		}
	}
}
