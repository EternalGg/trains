package isCousins

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Location struct {
	Level int
	Val   int
	Dad   int
}

type FindXY struct {
	level int
	x     *Location
	y     *Location
}

func isCousins(root *TreeNode, x int, y int) bool {
	f := new(FindXY)
	f.x = new(Location)
	f.x.Val = x
	f.y = new(Location)
	f.y.Val = y
	f.Find(root)
	if f.y.Level == f.x.Level && f.x.Dad != f.y.Dad {
		return true
	} else {
		return false
	}
}

func (f *FindXY) Find(root *TreeNode) {
	if root != nil {
		f.level++
		if root.Left != nil {
			if root.Left.Val == f.x.Val {
				f.x.Level = f.level
				f.x.Dad = root.Val
			}
			if root.Left.Val == f.y.Val {
				f.y.Level = f.level
				f.y.Dad = root.Val
			}
			f.Find(root.Left)
		}
		if root.Right != nil {
			if root.Right.Val == f.x.Val {
				f.x.Level = f.level
				f.x.Dad = root.Val
			}
			if root.Right.Val == f.y.Val {
				f.y.Level = f.level
				f.y.Dad = root.Val
			}
			f.Find(root.Right)
		}
		f.level--
	}
	return
}
