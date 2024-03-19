package levelOrder

type Node struct {
	Val      int
	Children []*Node
}

type Result struct {
	list [][]int
	high int
}

func levelOrder(root *Node) [][]int {
	r := new(Result)
	r.list = make([][]int, 0)
	r.high = 0
	r.levelOrderDigger(root)
	return r.list
}

func (r *Result) levelOrderDigger(root *Node) {
	if root != nil {

		if r.high >= len(r.list) {
			list := make([]int, 0)
			r.list = append(r.list, list)
		}
		r.list[r.high] = append(r.list[r.high], root.Val)

		for i := 0; i < len(root.Children); i++ {
			r.high++
			r.levelOrderDigger(root.Children[i])
		}
	}
	r.high--
}
