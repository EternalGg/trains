package pathSum

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Count struct {
	count  int
	list   []int
	result [][]int
	target int
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	c := new(Count)
	c.target = targetSum
	c.result = make([][]int, 0)
	c.list = make([]int, 0)
	c.count = 0
	c.cTree(root)
	return c.result
}

func (t *Count) cTree(root *TreeNode) {
	if root != nil {
		t.target -= root.Val
		t.list = append(t.list, root.Val)
		fmt.Println(t.list, t.target)
		if t.target == 0 && root.Left == nil && root.Right == nil {
			newList := make([]int, len(t.list))
			for i, num := range t.list {
				newList[i] = num
			}
			t.result = append(t.result, newList)
			fmt.Println(t.result)
		}
		t.cTree(root.Left)
		t.cTree(root.Right)
		t.list = t.list[:len(t.list)-1]
		t.target += root.Val
	}
	return
}
