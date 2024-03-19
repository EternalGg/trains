package lowestCommonAncestorr

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Count struct {
	tree   []*TreeNode
	list   []int
	result [][]int
}

func (t *Count) cTree(root, p, q *TreeNode) {
	if root != nil {
		t.tree = append(t.tree, root)
		t.list = append(t.list, root.Val)
		if root == p || root == q {
			list := make([]int, len(t.list))
			for i, num := range t.list {
				list[i] = num
			}
			t.result = append(t.result, list)
		}
		t.cTree(root.Left, p, q)
		t.cTree(root.Right, p, q)
		t.list = t.list[:len(t.list)-1]
	}
	return
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	c := new(Count)
	c.list = make([]int, 0)
	c.result = make([][]int, 0)
	c.cTree(root, p, q)

	iMap := make(map[int]bool)
	for _, i2 := range c.result[0] {
		iMap[i2] = true
	}
	tMap := make(map[int]*TreeNode)
	for _, node := range c.tree {
		if iMap[node.Val] {
			tMap[node.Val] = node
		}
	}

	for i := len(c.result[1]) - 1; i >= 0; i-- {
		if iMap[c.result[1][i]] {
			return tMap[c.result[1][i]]
		}
	}

	return nil
}
