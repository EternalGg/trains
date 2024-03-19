package postorder

type Node struct {
	Val      int
	Children []*Node
}

type Result struct {
	list []int
}

func postorder(root *Node) []int {
	r := new(Result)
	r.postorderDigger(root)
	return r.list
}

func (r *Result) postorderDigger(root *Node) {
	if root != nil {
		for i := 0; i < len(root.Children); i++ {
			r.postorderDigger(root.Children[i])
		}
		r.list = append(r.list, root.Val)
	}
}
