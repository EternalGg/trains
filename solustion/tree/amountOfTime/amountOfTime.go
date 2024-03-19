package amountOfTime

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type digger struct {
	max int
	now int
	pos int
}

func (d *digger) TreeSearch(root *TreeNode, target int) {
	if root != nil {
		d.now++
		if d.now > d.max {
			d.max = d.now
		}
		if root.Val == target {
			d.pos = d.now
		}
		d.TreeSearch(root.Left, target)
		d.TreeSearch(root.Left, target)
		d.now--
	}
	return
}

func amountOfTime(root *TreeNode, start int) int {
	if root.Right == nil && root.Left == nil {
		return 0
	}
	if start == root.Val {
		d := new(digger)
		d.TreeSearch(root, start)
		return d.max
	} else {
		ld := new(digger)
		ld.pos = -1
		rd := new(digger)
		rd.pos = -1

		ld.TreeSearch(root.Left, start)
		rd.TreeSearch(root.Right, start)
		max := 0
		fmt.Println(ld, rd)
		if ld.pos == -1 {
			cut := rd.max - rd.pos
			max = ld.max + rd.pos
			if max > cut {
				return max
			} else {
				return cut
			}
		} else {
			cut := ld.max - ld.pos
			max = rd.max + ld.pos
			if max > cut {
				return max
			} else {
				return cut
			}
		}
	}
}
