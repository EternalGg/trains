package connect

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

type LevelTree struct {
	levelLast map[int]*Node
	level     int
}

func levelTreeInit() LevelTree {
	l := new(LevelTree)
	l.levelLast = make(map[int]*Node)
	l.level = 0
	return *l
}

func connect(root *Node) *Node {
	l := levelTreeInit()
	l.levelConnect(root)
	return root
}

func (l LevelTree) levelConnect(root *Node) {
	if root != nil {
		if l.levelLast[l.level] == nil {
			l.levelLast[l.level] = root
			root.Next = nil
		} else {
			l.levelLast[l.level].Next = root
			l.levelLast[l.level] = root
			root.Next = nil
		}
		l.level++
		l.levelConnect(root.Left)
		l.levelConnect(root.Right)
		l.level--
	}
	return
}
