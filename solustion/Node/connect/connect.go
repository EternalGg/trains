package connect

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

type LevelTree struct {
	level int
	max   int
	lMap  map[int][]*Node
}

func LevelTreeInit() LevelTree {
	l := new(LevelTree)
	l.level = 0
	l.max = 0
	l.lMap = make(map[int][]*Node)
	return *l
}

func levelOrder(root *Node) [][]*Node {
	l := LevelTreeInit()

	result := make([][]*Node, 0)
	l.lorder(root)
	for i := 0; i < l.max; i++ {
		result = append(result, l.lMap[i])
	}

	return result
}

func (l *LevelTree) lorder(root *Node) {
	if root != nil {
		l.lMap[l.level] = append(l.lMap[l.level], root)
		l.level++
		if l.level > l.max {
			l.max = l.level
		}
		l.lorder(root.Left)
		l.lorder(root.Right)
		l.level--
	}
	return
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	list := levelOrder(root)
	for i := 0; i < len(list); i++ {

		for j := 0; j < len(list[i])-1; j++ {
			fmt.Println(list[i][j].Val)
			list[i][j].Next = list[i][j+1]
		}
	}
	return list[0][0]
}
