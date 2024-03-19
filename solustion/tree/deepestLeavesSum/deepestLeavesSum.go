package deepestLeavesSum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type LevelTree struct {
	level int
	max   int
	lMap  map[int][]int
}

func LevelTreeInit() LevelTree {
	l := new(LevelTree)
	l.level = 0
	l.max = 0
	l.lMap = make(map[int][]int)
	return *l
}

func levelOrder(root *TreeNode) [][]int {
	l := LevelTreeInit()

	result := make([][]int, 0)
	l.lorder(root)
	for i := 0; i < l.max; i++ {
		result = append(result, l.lMap[i])
	}
	return result
}

func (l *LevelTree) lorder(root *TreeNode) {
	if root != nil {
		l.lMap[l.level] = append(l.lMap[l.level], root.Val)
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

func deepestLeavesSum(root *TreeNode) int {
	list := levelOrder(root)
	length := len(list) - 1
	result := 0
	for _, i2 := range list[length] {
		result += i2
	}
	return result
}
