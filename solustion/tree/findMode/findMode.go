package findMode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Digger struct {
	iMap map[int]int
	Max  int
	List []int
}

func findMode(root *TreeNode) []int {
	d := new(Digger)
	d.iMap = make(map[int]int)
	d.treeDigger(root)
	return d.List
}

func (d *Digger) treeDigger(root *TreeNode) {
	if root == nil {
		return
	}
	d.iMap[root.Val]++
	if d.iMap[root.Val] >= d.Max {
		if d.iMap[root.Val] > d.Max {
			d.Max = d.iMap[root.Val]
			d.List = make([]int, 0)
		}
		d.List = append(d.List, root.Val)
	}

	d.treeDigger(root.Left)
	d.treeDigger(root.Right)

	return
}
