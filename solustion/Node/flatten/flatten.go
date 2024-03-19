package flatten

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}
